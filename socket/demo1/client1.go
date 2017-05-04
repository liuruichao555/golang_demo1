package main

import (
	"net"
	"fmt"
)

func main() {
	server, err := net.ResolveTCPAddr("tcp4", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		panic(err)
	}
	defer  conn.Close()

	conn.Write([]byte("bye"))
	buffer := make([]byte, 204)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buffer[0: n]))
}