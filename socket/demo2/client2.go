package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
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
	server, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入发送的内容：")
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		conn.Write(line)
		conn.Write([]byte("\n"))
		if "bye" == string(line) {
			fmt.Println("Bye Bye")
			break
		}
		// 读取服务器返回的信息
		data, err := readData(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println("client: " + data)
	}
}