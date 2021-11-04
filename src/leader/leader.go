package main

import (
	"context"
	"fmt"
	"go-squidgame/api/leaderpb"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

type GameState struct {
	stage      int32
	row        int32
	hasStarted bool
	moves      [16]int32
	hasMoved   [16]bool
	moveSums   [16]int32
	leaderMove int32
	hasLost    [16]bool
	numTeams   int
	teams      [16]int32
}

var state GameState

func main() {
	// Set initial values
	state.stage = 1
	state.row = 1
	state.hasStarted = false
	state.leaderMove = -1
	state.numTeams = -1
	for i := 0; i < 16; i++ {
		state.moves[i] = -1
		state.hasMoved[i] = false
		state.moveSums[i] = 0
		state.hasLost[i] = false
		state.teams[i] = -1
	}

	// Start server
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

func win(winners []int) {
	fmt.Println("¡El juego del calamar ha finalizado!")
	fmt.Println("Los ganadores son: ")

	for i := 0; i < len(winners); i++ {
		fmt.Println("Jugador ", winners[i])
	}
}

func lose(id int) bool {
	// Note: Player is informed in next GetPlayerState call

	state.hasLost[id] = true

	// Felipe: Llama a tu función aquí.

	fmt.Println("El jugador ", id+1, " ha muerto.")

	hasLostCount := 0
	winner := -1
	for i := 0; i < 16; i++ {
		if state.hasLost[i] {
			hasLostCount++
		} else {
			winner = i
		}
	}

	if hasLostCount == 15 {
		win([]int{winner})

		return true
	}

	return false
}

func nextStage(stage int32) {
	fmt.Println("¡Fin etapa ", stage, "!")
	fmt.Println("Los jugadores vivos son: ")

	// Reset values
	state.stage = stage + 1
	state.row = 1
	state.leaderMove = -1
	for i := 0; i < 16; i++ {
		if !state.hasLost[i] {
			fmt.Println("Jugador ", i+1)
		}

		state.moves[i] = -1
		state.hasMoved[i] = false
		state.moveSums[i] = 0
	}

	if state.stage == 2 {
		// TEAM UP
	} else if state.stage == 3 {
		// TEAM UP
	}

	// GET INPUT
}

func nextRow() {
	// Paula: Llama a Save aquí

	rand.Seed(time.Now().UnixNano())

	if state.stage == 1 {
		// Leader moves
		state.leaderMove = rand.Int31n(5) + 6

		// Players who move a number >= the leader's lose
		for i := 0; i < 16; i++ {
			if !state.hasLost[i] && state.moves[i] >= state.leaderMove {
				gameOver := lose(i)
				if gameOver {
					return
				}
			}
		}

		// If everyone sums at least 21, go to next stage
		allSumsAtLeast21 := true
		for i := 0; i < 16; i++ {
			if !state.hasLost[i] && state.moveSums[i] < 21 {
				allSumsAtLeast21 = false
				break
			}
		}

		if allSumsAtLeast21 {
			nextStage(1)
			return
		}

		// If in last row, players who sum < 21 lose, go to next stage
		if state.row == 4 {
			for i := 0; i < 16; i++ {
				if !state.hasLost[i] && state.moveSums[i] < 21 {
					gameOver := lose(i)
					if gameOver {
						return
					}
					break
				}
			}

			nextStage(1)
			return
		}
	} else if state.stage == 2 {
		// Leader moves
		state.leaderMove = rand.Int31n(4) + 1

		// Calculate leader and team parities
		leaderParity := state.leaderMove % 2

		team1Sum := int32(0)
		team2Sum := int32(0)
		for i := 0; i < 16; i++ {
			if !state.hasLost[i] && state.teams[i] == 1 {
				team1Sum += state.moves[i]
			}
			if !state.hasLost[i] && state.teams[i] == 2 {
				team2Sum += state.moves[i]
			}
		}

		team1Parity := team1Sum % 2
		team2Parity := team2Sum % 2

		// Handle all cases in truth table
		if team1Parity != leaderParity && team2Parity != leaderParity {
			fmt.Println("Ningún equipo ha obtenido la misma paridad que el líder.")

			// Leader moves
			state.leaderMove = rand.Int31n(2) + 1
			fmt.Println("El equipo ", state.leaderMove, " ha sido eliminado.")

			for i := 0; i < 16; i++ {
				if !state.hasLost[i] && state.teams[i] == state.leaderMove {
					gameOver := lose(i)
					if gameOver {
						return
					}
				}
			}
		} else if team1Parity == leaderParity && team2Parity != leaderParity {
			fmt.Println("El equipo 1 ha obtenido la misma paridad que el líder.")
			fmt.Println("El equipo 2 ha sido eliminado.")

			for i := 0; i < 16; i++ {
				if !state.hasLost[i] && state.teams[i] == 2 {
					gameOver := lose(i)
					if gameOver {
						return
					}
				}
			}
		} else if team1Parity != leaderParity && team2Parity == leaderParity {
			fmt.Println("El equipo 2 ha obtenido la misma paridad que el líder.")
			fmt.Println("El equipo 1 ha sido eliminado.")

			for i := 0; i < 16; i++ {
				if !state.hasLost[i] && state.teams[i] == 1 {
					gameOver := lose(i)
					if gameOver {
						return
					}
				}
			}
		}

		// Go to next stage
		nextStage(2)

		return
	} else if state.stage == 3 {
		// Leader moves
		state.leaderMove = rand.Int31n(10) + 1

		for i := 1; i <= state.numTeams; i++ {
			var playersInTeam []int

			for j := 0; j < 16; j++ {
				if !state.hasLost[j] && state.teams[j] == int32(i) {
					playersInTeam = append(playersInTeam, j)
				}
			}

			if state.moves[playersInTeam[0]] == state.moves[playersInTeam[1]] {
				win(playersInTeam)
			} else {
				if math.Abs(float64(state.moves[playersInTeam[0]]-state.leaderMove)) < math.Abs(float64(state.moves[playersInTeam[1]]-state.leaderMove)) {
					win([]int{playersInTeam[0]})
				} else {
					win([]int{playersInTeam[1]})
				}
			}

			return
		}
	}

	// Reset values
	state.leaderMove = -1
	for i := 0; i < 16; i++ {
		state.moves[i] = -1
		state.hasMoved[i] = false
		state.moveSums[i] = 0
	}

	state.row++
}

func (*server) GetPlayerState(ctx context.Context, req *leaderpb.GetPlayerStateRequest) (*leaderpb.GetPlayerStateResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	id := req.GetPlayerId()
	log.Println(id)

	stage := state.stage
	row := state.row
	hasStarted := state.hasStarted
	hasMoved := state.hasMoved[id]
	hasLost := state.hasLost[id]

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

	state.moves[id] = move
	state.hasMoved[id] = true
	state.moveSums[id] += move

	allHaveMoved := true
	for i := 0; i < 16; i++ {
		if !state.hasMoved[i] {
			allHaveMoved = false
			break
		}
	}

	if allHaveMoved {
		nextRow()
	}

	result := int32(1)

	res := &leaderpb.SendPlayerMoveResponse{
		Result: result,
	}
	return res, nil
}
