package main

import (
	"context"
	"fmt"
	"go-squidgame/api/poolpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50056")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	poolpb.RegisterPoolServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func (*server) GetPool(ctx context.Context, req *poolpb.GetPoolRequest) (*poolpb.GetPoolResponse, error) {
	// Unpack request
	request := req.GetRequest()

	fmt.Println(request)

	// Pack response
	pool := 9001

	// Send response
	res := &poolpb.GetPoolResponse{
		Pool: int32(pool),
	}
	return res, nil
}
