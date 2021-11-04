package main

import (
	"context"
	"fmt"
	"go-squidgame/api/namenodepb"
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
	namenodepb.RegisterNamenodeServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func (*server) Save(ctx context.Context, req *namenodepb.SaveRequest) (*namenodepb.SaveResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	moves := req.GetMoves()
	log.Println(moves)
	result := int32(1)
	res := &namenodepb.SaveResponse{
		Result: result,
	}
	return res, nil
}
