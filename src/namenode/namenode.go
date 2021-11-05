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

var cd1 datanodepb.DatanodeServiceClient
var cd2 datanodepb.DatanodeServiceClient
var cd3 datanodepb.DatanodeServiceClient

func main() {
	pdh.stage1 = [16]string{}
	pdh.stage2 = [16]string{}
	pdh.stage3 = [16]string{}
	global_stage.start = true
	global_stage.actual_stage = int32(1)

	// Connect to datanodeone server
	fmt.Println("Starting Client...")
	cc, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cd1 = datanodepb.NewDatanodeServiceClient(cc)

	// Connect to datanodetwo server
	fmt.Println("Starting Client...")
	cc, err = grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cd2 = datanodepb.NewDatanodeServiceClient(cc)

	// Connect to datanodethree server
	fmt.Println("Starting Client...")
	cc, err = grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()
	cd3 = datanodepb.NewDatanodeServiceClient(cc)

	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "0.0.0.0:50052")
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
	moves := req.GetMoves()
	row := req.GetRow()
	stage := req.GetStage()

	/*fmt.Println("[DEBUG]")
	fmt.Println("Moves ", moves)
	fmt.Println("Ronda ", row)
	fmt.Println("Etapa ", stage)*/

	//De este modo se asignan los datanodes solo en la primera ronda y al cambio de etapa
	if global_stage.start {
		mapPlayersToDatanodes(moves, stage)
		global_stage.start = false
		//]ln("Asignación primera ronda")
	} else {
		if stage > global_stage.actual_stage {
			mapPlayersToDatanodes(moves, stage)
			global_stage.actual_stage = stage
			//]ln("Asignación ronda", global_stage.actual_stage)
		}
	}
	updateMoves(moves)

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

			// Send request
			_, err := cd1.Write(context.Background(), datanode_req)
			if err != nil {
				log.Fatalf("Error Call RPC: %v", err)
			}
		case 1:
			datanode_req := &datanodepb.WriteRequest{
				Stage:   stage,
				Moves:   dn_moves.datanode2,
				Players: players.datanode2,
			}

			// Send request
			_, err := cd2.Write(context.Background(), datanode_req)
			if err != nil {
				log.Fatalf("Error Call RPC: %v", err)
			}
		case 2:
			datanode_req := &datanodepb.WriteRequest{
				Stage:   stage,
				Moves:   dn_moves.datanode3,
				Players: players.datanode3,
			}

			// Send request
			_, err := cd3.Write(context.Background(), datanode_req)
			if err != nil {
				log.Fatalf("Error Call RPC: %v", err)
			}
		}

	}

	//]ln(moves)
	result := int32(1)
	res := &namenodepb.SaveResponse{
		Result: result,
	}
	return res, nil
}

func (*server) Open(ctx context.Context, req *namenodepb.OpenRequest) (*namenodepb.OpenResponse, error) {
	log.Printf("Greet was invoked  with %v\n", req)
	stage := req.GetStage()
	player := req.GetPlayer()
	var moves_stage1 []int32
	move_stage2 := int32(-1)
	move_stage3 := int32(-1)
	datanode_req := &datanodepb.ReadRequest{
		Stage:  stage,
		Player: player,
	}
	var res *datanodepb.ReadResponse
	var err error
	if int32InSlice(player, players.datanode1) {
		// Send request
		res, err = cd1.Read(context.Background(), datanode_req)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}
	} else if int32InSlice(player, players.datanode2) {
		// Send request
		res, err = cd2.Read(context.Background(), datanode_req)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}
	} else {
		// Send requestmapPlayersToDatanodes(moves []int32, stage int32) {
		res, err = cd3.Read(context.Background(), datanode_req)
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}
	}
	moves_stage1 = res.GetMovesStage1()
	move_stage2 = res.GetMoveStage2()
	move_stage3 = res.GetMoveStage3()
	response := &namenodepb.OpenResponse{
		MovesStage1: moves_stage1[:],
		MoveStage2:  move_stage2,
		MoveStage3:  move_stage3,
	}
	return response, nil
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
}
func mapPlayersToDatanodes(moves []int32, stage int32) {
	dn_moves.datanode1 = []int32{}
	dn_moves.datanode2 = []int32{}
	dn_moves.datanode3 = []int32{}
	players.datanode1 = []int32{}
	players.datanode2 = []int32{}
	players.datanode3 = []int32{}
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
	//fmt.Println(players, dn_moves)
}
func updateMoves(moves []int32) {
	dn_moves.datanode1 = []int32{}
	dn_moves.datanode2 = []int32{}
	dn_moves.datanode3 = []int32{}

	for i := 0; i < len(players.datanode1); i++ {
		dn_moves.datanode1 = append(dn_moves.datanode1, moves[players.datanode1[i]])
	}
	for i := 0; i < len(players.datanode2); i++ {
		dn_moves.datanode2 = append(dn_moves.datanode2, moves[players.datanode2[i]])
	}
	for i := 0; i < len(players.datanode3); i++ {
		dn_moves.datanode3 = append(dn_moves.datanode3, moves[players.datanode3[i]])
	}
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
func int32InSlice(a int32, list []int32) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
