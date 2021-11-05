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
var s *grpc.Server

func main() {
	// Set initial values
	state.stage = 1
	state.row = 1
	state.hasStarted = true
	state.leaderMove = -1
	state.numTeams = -1
	for i := 0; i < 16; i++ {
		state.moves[i] = -1
		state.hasMoved[i] = false
		state.moveSums[i] = 0
		state.hasLost[i] = false
		state.teams[i] = -1
	}

	// Welcome
	fmt.Println("¡Bienvenido al Juego del Calamar!")
	fmt.Println("Al finalizar cada etapa podrá preguntar por las jugadas de un jugador o")
	fmt.Println("pasar a la siguiente etapa.")

	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s = grpc.NewServer()
	leaderpb.RegisterLeaderServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func win(winners []int) {
	fmt.Println("¡El juego del calamar ha finalizado!")
	fmt.Println("Los ganadores son: ")

	for i := 0; i < len(winners); i++ {
		fmt.Println("Jugador ", winners[i]+1)
	}

	time.Sleep(time.Duration(1<<63 - 1))
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

	// Reset values and keep track of players who have not lost
	state.stage = stage + 1
	state.row = 1
	state.hasStarted = false
	state.leaderMove = -1

	hasNotLostCount := 0
	var hasNotLost []int
	for i := 0; i < 16; i++ {
		if !state.hasLost[i] {
			fmt.Println("Jugador ", i+1)
			hasNotLost = append(hasNotLost, i)
			hasNotLostCount++
		}

		state.moves[i] = -1
		state.hasMoved[i] = false
		state.moveSums[i] = 0
		state.teams[i] = -1
	}

	// Ask for input
	for {
		// Paula: El menú por si se quiere llamar a Open irá aquí

		fmt.Println("¿Desea pasar a la siguiente etapa? (Y/N)")

		var input string
		fmt.Scanln(&input)
		if input == "Y" {
			state.hasStarted = true
			break
		}
	}

	rand.Seed(time.Now().UnixNano())

	if state.stage == 2 {
		// If number of players is odd, a random player loses
		if hasNotLostCount%2 == 1 {
			toLose := rand.Intn(hasNotLostCount)
			gameOver := lose(hasNotLost[toLose])
			if gameOver {
				return
			}
			removeIndex(hasNotLost, toLose)
			hasNotLostCount--
		}

		// Team up in two teams
		for i := 0; i < hasNotLostCount; i++ {
			if i%2 == 0 {
				state.teams[hasNotLost[i]] = 1
			} else {
				state.teams[hasNotLost[i]] = 2
			}
		}
	} else if state.stage == 3 {
		// If number of players is odd, a random player loses
		if hasNotLostCount%2 == 1 {
			toLose := rand.Intn(hasNotLostCount)
			gameOver := lose(hasNotLost[toLose])
			if gameOver {
				return
			}
			hasNotLostCount--
			removeIndex(hasNotLost, toLose)
		}

		// Team up in pairs
		for i := 0; i < hasNotLostCount; i++ {
			if 2*i >= hasNotLostCount {
				state.numTeams = i
				break
			}
			state.teams[hasNotLost[2*i]] = int32(i + 1)
			state.teams[hasNotLost[2*i+1]] = int32(i + 1)
		}
	}
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
		} else {
			fmt.Println("Ambos equipos han obtenido la misma paridad que el líder.")
		}

		// Go to next stage
		nextStage(2)

		return
	} else if state.stage == 3 {
		// Leader moves
		state.leaderMove = rand.Int31n(10) + 1

		var winners []int
		for i := 1; i <= state.numTeams; i++ {
			var playersInTeam []int

			for j := 0; j < 16; j++ {
				if !state.hasLost[j] && state.teams[j] == int32(i) {
					playersInTeam = append(playersInTeam, j)
				}
			}

			//fmt.Println("[DEBUG] ", i, playersInTeam)

			if state.moves[playersInTeam[0]] == state.moves[playersInTeam[1]] {
				winners = append(winners, playersInTeam...)
			} else {
				if math.Abs(float64(state.moves[playersInTeam[0]]-state.leaderMove)) < math.Abs(float64(state.moves[playersInTeam[1]]-state.leaderMove)) {
					winners = append(winners, playersInTeam[0])
				} else {
					winners = append(winners, playersInTeam[1])
				}
			}
		}

		win(winners)

		return
	}

	// Reset values
	state.leaderMove = -1
	for i := 0; i < 16; i++ {
		state.moves[i] = -1
		state.hasMoved[i] = false
	}

	state.row++
}

func (*server) GetPlayerState(ctx context.Context, req *leaderpb.GetPlayerStateRequest) (*leaderpb.GetPlayerStateResponse, error) {
	id := req.GetPlayerId()

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
	id := req.GetPlayerId()
	move := req.GetMove()

	//fmt.Println("[DEBUG] El jugador ", id+1, " en la ronda ", state.row, " ha enviado un ", move)

	state.moves[id] = move
	state.hasMoved[id] = true
	state.moveSums[id] += move

	allHaveMoved := true
	for i := 0; i < 16; i++ {
		if !state.hasLost[i] && !state.hasMoved[i] {
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
