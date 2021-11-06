package main

import (
	"bufio"
	"context"
	"fmt"
	"go-squidgame/api/datanodepb"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50054")
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
	log.Printf("Greet was invoked  with %v\n", req)
	stage := req.GetStage()
	player := req.GetPlayer()
	moves_stage1 := []int32{-1, -1, -1, -1, -1, -1}
	move_stage2 := int32(-1)
	move_stage3 := int32(-1)
	// Lee la información del jugador por cada Etapa
	for i := 1; i <= int(stage); i++ {
		data := readData(int32(i), player)
		if i == 1 {
			moves_stage1 = data
		} else if i == 2 {
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
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
	filename := "jugador_" + fmt.Sprint(player+1) + "__etapa_" + fmt.Sprint(stage) + ".txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
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
