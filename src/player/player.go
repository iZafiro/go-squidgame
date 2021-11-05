package main

import (
	"context"
	"fmt"
	"go-squidgame/api/leaderpb"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

var exited = [15]bool{
	false, false, false, false, false,
	false, false, false, false, false,
	false, false, false, false, false,
}

func main() {
	// Connect to server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := leaderpb.NewLeaderServiceClient(cc)

	// Welcome
	fmt.Println("¡Bienvenido al Juego del Calamar!")

	// Ask to join
	for {
		fmt.Println("¿Desea iniciar el juego? (Y/N)")

		var input string
		fmt.Scanln(&input)
		if input == "Y" {
			break
		}
	}

	// Create bots
	for i := 1; i < 16; i++ {
		go bot(i, c)
	}

	// Human player loop
	for {
		stage, row, hasStarted, hasMoved, hasLost := getState(0, c)

		time.Sleep(200 * time.Millisecond)

		if hasLost {
			break
		}

		if hasStarted && !hasMoved {
			fmt.Println("Escriba su jugada (etapa ", stage, ", ronda ", row, "): ")

			var input int32
			fmt.Scanln(&input)

			sendMove(0, input, c)
			time.Sleep(200 * time.Millisecond)
		}
	}

	for {
		fmt.Println("¿Desea interrumpir o finalizar el juego? (Y/N)")

		var input string
		fmt.Scanln(&input)
		if input == "Y" {
			break
		}
	}
}

func bot(id int, c leaderpb.LeaderServiceClient) {
	// Bot player loop
	for {
		stage, _, hasStarted, hasMoved, hasLost := getState(id, c)
		time.Sleep(200 * time.Millisecond)

		if hasLost {
			exited[id-1] = true
			return
		}

		if hasStarted && !hasMoved {
			rand.Seed(time.Now().UnixNano())

			var move int32
			if stage == 1 {
				move = rand.Int31n(10) + 1
			} else if stage == 2 {
				move = rand.Int31n(4) + 1
			} else if stage == 3 {
				move = rand.Int31n(10) + 1
			}

			//fmt.Println("[DEBUG] El jugador ", id+1, " en la ronda ", row, " ha enviado un ", move)

			sendMove(id, move, c)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func getState(id int, c leaderpb.LeaderServiceClient) (int32, int32, bool, bool, bool) {
	req := &leaderpb.GetPlayerStateRequest{
		PlayerId: int32(id),
	}

	res, err := c.GetPlayerState(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}
	return res.Stage, res.Row, res.HasStarted, res.HasMoved, res.HasLost
}

func sendMove(id int, move int32, c leaderpb.LeaderServiceClient) int32 {
	req := &leaderpb.SendPlayerMoveRequest{
		PlayerId: int32(id),
		Move:     move,
	}

	res, err := c.SendPlayerMove(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}
	return res.Result
}
