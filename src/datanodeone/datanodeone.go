package main

import (
	"bufio"
	"context"
	"fmt"
	"go-squidgame/api/datanodepb"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	folder := "src/datanodeone/out"

	RemoveContents(folder)

	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	datanodepb.RegisterDatanodeServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}

}

// Write: Recibe la petición por parte del namenode para almacenar la información de una ronda
//		(de algunos jugadores) en archivos de texto por jugador y etapa
func (*server) Write(ctx context.Context, req *datanodepb.WriteRequest) (*datanodepb.WriteResponse, error) {
	moves := req.GetMoves()
	stage := req.GetStage()
	players := req.GetPlayers()

	// Almacena la información en archivos de texto por jugador
	for i := 0; i < len(players); i++ {
		saveData(moves[i], stage, players[i])
	}
	result := int32(1)

	// Retorna que se almacenó con éxito
	res := &datanodepb.WriteResponse{
		Result: result,
	}
	return res, nil
}

// Read: Recibe la petición del namenode para leer la información de un jugador
func (*server) Read(ctx context.Context, req *datanodepb.ReadRequest) (*datanodepb.ReadResponse, error) {
	stage := req.GetStage()
	player := req.GetPlayer()
	moves_stage1 := []int32{}
	move_stage2 := int32(-1)
	move_stage3 := int32(-1)
	// Lee la información del jugador en la Etapa correspondiente
	data := readData(int32(stage), player)
	if stage == 1 {
		moves_stage1 = data
	} else if stage == 2 {
		if len(data) > 0 {
			move_stage2 = data[0]
		} else {
			move_stage2 = -1
		}
	} else {
		if len(data) > 0 {
			move_stage3 = data[0]
		} else {
			move_stage3 = -1
		}

	}

	// Retorna la información del jugador por cada Etapa
	res := &datanodepb.ReadResponse{
		MovesStage1: moves_stage1,
		MoveStage2:  move_stage2,
		MoveStage3:  move_stage3,
	}
	return res, nil
}

// Función para guardar la información de un jugador en una ronda en una Etapa determinada
func saveData(move int32, stage int32, player int32) {
	filename := "jugador_" + fmt.Sprint(player+1) + "__etapa_" + fmt.Sprint(stage) + ".txt"
	folder := "src/datanodeone/out/"
	f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	if move != -1 {
		fmt.Fprintln(f, move)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Función para leer la información de los movimientos de un jugador en una Etapa
func readData(stage int32, player int32) []int32 {
	moves_response := []int32{}
	folder := "src/datanodeone/out/"
	filename := "jugador_" + fmt.Sprint(player+1) + "__etapa_" + fmt.Sprint(stage) + ".txt"
	file, err := os.Open(folder + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		result := int32(i)
		moves_response = append(moves_response, result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return moves_response
}

// Borra los archivos creados en la partida
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
