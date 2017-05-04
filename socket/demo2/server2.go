package main

import (
	"fmt"
	"net"
	"io"
	"strings"
)

func readData(conn net.Conn) (string, error) {
	buf := make([]byte, 2048)
	var data string
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		data += string(buf[0: n])
		if strings.HasSuffix(data, "\n") {
			break
		}
	}
	data = strings.Replace(data, "\n", "", 1)
	return data , nil
}

func main() {
	//addr := "localhost"
	//port := 8080
	netListener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer netListener.Close()
	fmt.Println("server started, listen port: ", 8080)

	for {
		conn, err := netListener.Accept()
		if err != nil {
			fmt.Errorf("conn error: %s.\n", err.Error())
			conn.Close()
			continue
		}
		fmt.Println("connection remote addr: " + conn.RemoteAddr().String())

		for {
			// 读取结果
			data, err := readData(conn)
			if err != nil {
				panic(err)
			}

			fmt.Println("receive data: ", data)
			if "bye" == data {
				break
			}
			conn.Write([]byte("呵呵哒"))
			conn.Write([]byte("\n"))
		}
		fmt.Println("process done.", conn.RemoteAddr().String())
	}
}