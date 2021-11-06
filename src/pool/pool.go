package main

import (
	"bufio"
	"context"
	"fmt"
	"go-squidgame/api/poolpb"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	//Connectar a RabbitMQ
	conn, err := amqp.Dial("amqp://usuario2:pass2@10.6.43.59:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Succesfully connected to RabbitMQ instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
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

	//consume los mensajes de la cola
	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	add_to_pool("")

	//Abrir thread del listener de RabbitMQ
	go func() {
		for d := range msgs {
			text := strings.Split(string(d.Body), " ")
			num, round := text[0], text[1]
			if num != "0" {
				amount := pool_total() + 100000000
				s := "Jugador_" + num + " Ronda_" + round + " " + strconv.Itoa(amount) + "\n"
				add_to_pool(s) // agrega jugador al pool
			} else {
				err := os.Remove("pool.txt") // elimina txt cuando recibe un 0
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()

	// Start server
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", "10.6.43.59:50056")
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
	pool := pool_total()
	// Send response
	res := &poolpb.GetPoolResponse{
		Pool: int32(pool),
	}
	return res, nil
}

func add_to_pool(s string) {

	//Crea o abre archivo txt
	f, err := os.OpenFile("pool.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//Escribe mensaje al final
	if _, err := f.Write([]byte(s)); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func pool_total() int {

	total := 0
	f, err := os.Open("pool.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		num, err := strconv.Atoi(text[2])
		if err != nil {
			log.Fatal(err)
		}
		total = num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//retorna el último monto
	return total
}
