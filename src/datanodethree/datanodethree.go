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
	l, err := net.Listen("tcp", "0.0.0.0:50055")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	datanodepb.RegisterDatanodeServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}

}
func (*server) Write(ctx context.Context, req *datanodepb.WriteRequest) (*datanodepb.WriteResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	moves := req.GetMoves()
	stage := req.GetStage()
	players := req.GetPlayers()
	for i := 0; i < len(players); i++ {
		saveData(moves[i], stage, players[i])
	}
	result := int32(1)
	res := &datanodepb.WriteResponse{
		Result: result,
	}
	return res, nil
}
func (*server) Read(ctx context.Context, req *datanodepb.ReadRequest) (*datanodepb.ReadResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	stage := req.GetStage()
	player := req.GetPlayer()
	moves_stage1 := [6]int32{}
	move_stage2 := int32(-1)
	move_stage3 := int32(-1)
	for i := 1; i <= int(stage); i++ {
		if i == 1 {
			moves_stage1 = readData(int32(i), player)
		} else if i == 2 {
			move_stage2 = readData(int32(i), player)[0]
		} else {
			move_stage3 = readData(int32(i), player)[0]
		}
	}

	res := &datanodepb.ReadResponse{
		MovesStage1: moves_stage1[:],
		MoveStage2:  move_stage2,
		MoveStage3:  move_stage3,
	}
	return res, nil
}
func saveData(move int32, stage int32, player int32) {
	filename := "jugador_" + fmt.Sprint(player) + "__ronda_" + fmt.Sprint(stage) + ".txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(f, move)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}
func readData(stage int32, player int32) [6]int32 {
	moves_response := [6]int32{-1, -1, -1, -1, -1, -1}
	filename := "jugador_" + fmt.Sprint(player) + "__ronda_" + fmt.Sprint(stage) + ".txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	index := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			panic(err)
		}
		result := int32(i)
		moves_response[index] = result
		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return moves_response
}
