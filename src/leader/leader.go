package main

import (
	"context"
	"fmt"
	"go-squidgame/api/leaderpb"
	"go-squidgame/api/namenodepb"
	"go-squidgame/api/poolpb"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type server struct{}

// This struct stores the full game state at a given row
// Player state is requested every 200 ms by players who have not lost
// and the leader responds with all relevant data
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
var cn namenodepb.NamenodeServiceClient

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

	// Connect to namenode server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cn = namenodepb.NewNamenodeServiceClient(cc)

	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50060")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s = grpc.NewServer()
	leaderpb.RegisterLeaderServiceServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}

// Helper function to remove entry at index from integer array
func removeIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func win(winners []int) {
	// Game over message
	fmt.Println("¡El juego del calamar ha finalizado!")

	// List winners
	fmt.Println("Los ganadores son: ")

	for i := 0; i < len(winners); i++ {
		fmt.Println("Jugador ", winners[i]+1)
	}

	time.Sleep(time.Duration(1<<63 - 1))
}

func lose(id int) bool {
	// Note: Player is informed in their next GetPlayerState call

	state.hasLost[id] = true
	str := strconv.Itoa(id) + " " + strconv.Itoa(int(state.row)) // "[num jugador] [num ronda]"
	addToQueue(str)                                              // agrega a la cola

	// Player game over message
	fmt.Println("El jugador ", id+1, " ha muerto.")

	// If only one player left, they win
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
	// End of stage message
	fmt.Println("¡Fin etapa ", stage, "!")

	// List players left
	fmt.Println("Los jugadores vivos son: ")

	// Reset values
	state.stage = stage + 1
	state.row = 1
	state.hasStarted = false
	state.leaderMove = -1

	// Keep track of players left
	hasNotLostCount := 0
	var hasNotLost []int
	for i := 0; i < 16; i++ {
		if !state.hasLost[i] {
			fmt.Println("Jugador ", i+1)
			hasNotLost = append(hasNotLost, i)
			hasNotLostCount++
		}

		// Reset values
		state.moves[i] = -1
		state.hasMoved[i] = false
		state.moveSums[i] = 0
		state.teams[i] = -1
	}

	// Ask for input
	for {
		fmt.Println("¿Qué desea hacer?")
		fmt.Println("1: Dar comienzo a la siguiente etapa.")
		fmt.Println("2: Preguntar por todas las jugadas de un jugador.")

		var input int
		fmt.Scanln(&input)
		if input == 1 {
			state.hasStarted = true
			break
		} else if input == 2 {
			var inputPlayer int32
			fmt.Println("Ingrese el número del jugador (1-16, 1 es el jugador humano): ")
			fmt.Scanln(&inputPlayer)

			movesStage1, moveStage2, moveStage3 := open(inputPlayer - 1)

			for i := 0; i < len(movesStage1); i++ {
				fmt.Println("Etapa 1")
				fmt.Println("En la ronda ", i+1, " jugó ", movesStage1[i])
			}

			if moveStage2 != -1 {
				fmt.Println("En la etapa 2 jugó ", moveStage2)
			}

			if moveStage3 != -1 {
				fmt.Println("En la etapa 3 jugó ", moveStage3)
			}
		}
	}

	rand.Seed(time.Now().UnixNano())

	// Beginning of stage logic
	// This includes restoring the parity of the number of players
	// and teaming up
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
	// End of row logic
	// This includes most winning / losing logic

	// Send save request to namenode
	save()

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

		// Handle all cases in truth table given by stage logic
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

		// Get winners in each team
		var winners []int
		for i := 1; i <= state.numTeams; i++ {

			// Get players in team
			var playersInTeam []int
			for j := 0; j < 16; j++ {
				if !state.hasLost[j] && state.teams[j] == int32(i) {
					playersInTeam = append(playersInTeam, j)
				}
			}

			// Stage logic
			if state.moves[playersInTeam[0]] == state.moves[playersInTeam[1]] {
				winners = append(winners, playersInTeam...)
			} else {
				if math.Abs(float64(state.moves[playersInTeam[0]]-state.leaderMove)) < math.Abs(float64(state.moves[playersInTeam[1]]-state.leaderMove)) {
					winners = append(winners, playersInTeam[0])
					gameOver := lose(playersInTeam[1])
					if gameOver {
						return
					}
				} else {
					winners = append(winners, playersInTeam[1])
					gameOver := lose(playersInTeam[0])
					if gameOver {
						return
					}
				}
			}
		}

		// Final call to win
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
	// Unpack request
	id := req.GetPlayerId()

	// Pack response
	stage := state.stage
	row := state.row
	hasStarted := state.hasStarted
	hasMoved := state.hasMoved[id]
	hasLost := state.hasLost[id]

	// Send response
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
	// Unpack request
	id := req.GetPlayerId()
	move := req.GetMove()

	// Update game state
	state.moves[id] = move
	state.hasMoved[id] = true
	state.moveSums[id] += move

	// If all players have moved in current row, do end of row logic
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

	// Pack response
	result := int32(1)

	// Send response
	res := &leaderpb.SendPlayerMoveResponse{
		Result: result,
	}
	return res, nil
}

func (*server) PlayerGetPool(ctx context.Context, req *leaderpb.PlayerGetPoolRequest) (*leaderpb.PlayerGetPoolResponse, error) {
	// Unpack request
	request := req.GetRequest()

	fmt.Println(request)

	// Connect to server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	c := poolpb.NewPoolServiceClient(cc)

	// Pack response
	pool := getPool(c)

	// Send response
	res := &leaderpb.PlayerGetPoolResponse{
		Pool: int32(pool),
	}
	return res, nil
}

func getPool(c poolpb.PoolServiceClient) int32 {
	// Pack request
	req := &poolpb.GetPoolRequest{
		Request: 1,
	}

	// Send request
	res, err := c.GetPool(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}
	return res.Pool
}

func addToQueue(str string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(str),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func save() int32 {
	var movesToSave []int32
	movesToSave = state.moves[:]

	// Pack request
	req := &namenodepb.SaveRequest{
		Stage: state.stage,
		Row:   state.row,
		Moves: movesToSave,
	}

	// Send request
	res, err := cn.Save(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}
	return res.Result
}

func open(p int32) ([]int32, int32, int32) {
	// Pack request
	req := &namenodepb.OpenRequest{
		Stage:  state.stage - 1,
		Player: p,
	}

	// Send request
	res, err := cn.Open(context.Background(), req)
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}
	return res.MovesStage1, res.MoveStage2, res.MoveStage3
}
