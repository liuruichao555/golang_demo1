package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
)

func main() {
	f, err := os.Open("/Users/liuruichao/shell/test/test4.sh")
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	for {
		//line, err := buf.ReadString('\n')
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(string(line))
	}
}
