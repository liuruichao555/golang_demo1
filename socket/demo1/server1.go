package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer netListen.Close()
	fmt.Println("server started...")

	isRun := true
	for {
		if !isRun {
			break
		}
		conn, err := netListen.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("address: %s.\n", conn.RemoteAddr())

		buffer := make([]byte, 2048)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			data := string(buffer[0: n])
			if data == "bye" {
				isRun = false
				conn.Write([]byte("bye"))
				break
			}
			fmt.Println("receive data: ", string(buffer[0: n]))
			conn.Write([]byte("hello"))
		}
	}
}