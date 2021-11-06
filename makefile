build:
	go build -o src/player/player.go
	go build -o src/leader/leader.go
	go build -o src/pool/pool.go
	go build -o src/namenode/namenode.go
	go build -o src/datanode1/datanode1.go
	go build -o src/datanode2/datanode2.go
	go build -o src/datanode3/datanode3.go

player:
	go run src/player/player.go

leader:
	go run src/leader/leader.go

pool:
	go run src/pool/pool.go

namenode:
	go run src/namenode/namenode.go

datanode1:
	go run src/datanode1/datanode1.go

datanode2:
	go run src/datanode2/datanode2.go

datanode3:
	go run src/datanode3/datanode3.go