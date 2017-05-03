package main

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	f, err := os.Open("/Users/liuruichao/shell/test/test4.sh")
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				continue
			}
			panic(err)
		}
		line = strings.Replace(line, "\n", "", 1)
		fmt.Println(line)
	}
}
