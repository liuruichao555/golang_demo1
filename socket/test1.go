package main

import (
	"net"
	"fmt"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer netListen.Close()
	fmt.Println("server started...")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("address: %s.\n", conn.RemoteAddr())

		buffer := make([]byte, 2048)
		for {
			data, err := conn.Read(buffer)
			if err != nil {
				panic(err)
			}
			fmt.Println("receive data: ", string(data))
			conn.Write([]byte("hello"))
		}
	}
}