package main

import (
	"net"
	"log"
)

const (
	port = ":8888"
)

type server struct {
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
