package main

import (
	"context"
	"fmt"
	"go-squidgame/api/leaderpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	leaderpb.RegisterLeaderServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func (*server) GetPlayerState(ctx context.Context, req *leaderpb.GetPlayerStateRequest) (*leaderpb.GetPlayerStateResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	id := req.GetPlayerId()
	log.Println(id)

	stage := int32(1)
	row := int32(1)
	hasStarted := true
	hasMoved := false
	hasLost := false

	res := &leaderpb.GetPlayerStateResponse{
		Stage:      stage,
		Row:        row,
		HasStarted: hasStarted,
		HasMoved:   hasMoved,
		HasLost:    hasLost,
	}
	return res, nil
}

func (*server) SendPlayerMove(ctx context.Context, req *leaderpb.SendPlayerMoveRequest) (*leaderpb.SendPlayerMoveResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	id := req.GetPlayerId()
	move := req.GetMove()
	log.Println(id, move)

	result := int32(1)

	res := &leaderpb.SendPlayerMoveResponse{
		Result: result,
	}
	return res, nil
}
