package main

import (
	"context"
	"fmt"
	"go-squidgame/api/datanodepb"
	"go-squidgame/api/namenodepb"
	"log"
	"math/rand"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

type players_datanodes_hash struct {
	stage1 [16]string
	stage2 [16]string
	stage3 [16]string
}

type players_datanodes struct {
	datanode1 []int32
	datanode2 []int32
	datanode3 []int32
}

type moves_datanodes struct {
	datanode1 []int32
	datanode2 []int32
	datanode3 []int32
}

type stages_state struct {
	start        bool
	actual_stage int32
}

var pdh players_datanodes_hash
var players players_datanodes
var dn_moves moves_datanodes
var global_stage stages_state

func main() {
	pdh.stage1 = [16]string{}
	pdh.stage2 = [16]string{}
	pdh.stage3 = [16]string{}
	global_stage.start = true
	global_stage.actual_stage = int32(1)
	mapPlayersToDatanodes([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 1)
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
	row := req.GetRow()
	stage := req.GetStage()

	//De este modo se asignan los datanodes solo en la primera ronda y al cambio de etapa
	if global_stage.start {
		mapPlayersToDatanodes(moves, stage)
		log.Println("Asignación primera ronda")
	} else {
		if stage > global_stage.actual_stage {
			mapPlayersToDatanodes(moves, stage)
			global_stage.actual_stage = stage
			log.Println("Asignación ronda", global_stage.actual_stage)
		}
	}

	for i := 0; i < len(moves); i++ {
		if stage == 1 {
			saveData(int32(i+1), row, pdh.stage1[i])
		} else if stage == 2 {
			saveData(int32(i+1), row, pdh.stage2[i])
		} else {
			saveData(int32(i+1), row, pdh.stage3[i])
		}

	}
	for node := 0; node < 3; node++ {
		switch node {
		case 0:
			datanode_req := &datanodepb.WriteRequest{
				Stage:   stage,
				Moves:   dn_moves.datanode1,
				Players: players.datanode1,
			}
		case 1:
			datanode_req := &datanodepb.WriteRequest{
				Stage:   stage,
				Moves:   dn_moves.datanode2,
				Players: players.datanode2,
			}
		case 2:
			datanode_req := &datanodepb.WriteRequest{
				Stage:   stage,
				Moves:   dn_moves.datanode3,
				Players: players.datanode3,
			}
		}

	}

	log.Println(moves)
	result := int32(1)
	res := &namenodepb.SaveResponse{
		Result: result,
	}
	return res, nil
}

func (*server) Open(ctx context.Context, req *namenodepb.OpenRequest) (*namenodepb.OpenResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	//stage := req.GetStage()
	//player := req.GetPlayer()
	moves_stage1 := [4]int32{-1, -1, -1, -1}
	move_stage2 := int32(-1)
	move_stage3 := int32(-1)
	res := &namenodepb.OpenResponse{
		MovesStage1: moves_stage1[:],
		MoveStage2:  move_stage2,
		MoveStage3:  move_stage3,
	}
	return res, nil
}

func saveData(player int32, row int32, ip string) {
	filename := "players_datanodes_hash.txt"
	text := "Jugador_" + fmt.Sprint(player) + " Ronda_" + fmt.Sprint(row) + " " + ip
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(f, text)
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
func mapPlayersToDatanodes(moves []int32, stage int32) {
	in := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	datanode_index := int(0)
	datanodes := [3]string{"0.0.0.0:50053", "0.0.0.0:50054", "0.0.0.0:50055"}
	for i := 0; i < 16; i++ {
		if i == 6 || i == 11 {
			datanode_index++
		}
		randomIndex := rand.Intn(len(in))
		if stage == 1 {
			pdh.stage1[in[randomIndex]] = datanodes[datanode_index]

		} else if stage == 2 {
			pdh.stage2[in[randomIndex]] = datanodes[datanode_index]

		} else {
			pdh.stage3[in[randomIndex]] = datanodes[datanode_index]

		}
		switch datanode_index {
		case 0:
			players.datanode1 = append(players.datanode1, int32(in[randomIndex]))
			dn_moves.datanode1 = append(dn_moves.datanode1, moves[in[randomIndex]])
		case 1:
			players.datanode2 = append(players.datanode2, int32(in[randomIndex]))
			dn_moves.datanode2 = append(dn_moves.datanode2, moves[in[randomIndex]])
		case 2:
			players.datanode3 = append(players.datanode3, int32(in[randomIndex]))
			dn_moves.datanode3 = append(dn_moves.datanode3, moves[in[randomIndex]])
		}
		in = remove(in, randomIndex)
	}
	fmt.Println(players, dn_moves)
}
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
