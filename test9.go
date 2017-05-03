package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("请输入名称简写：")
		databyte, _, _ := reader.ReadLine()
		data := string(databyte)
		if data == "bye" {
			fmt.Println("Bye Bye")
			break
		}
		switch data {
		case "lrc":
			fmt.Println("刘瑞超")
			break
		case "bzd":
			fmt.Println("不知道")
			break
		case "hhd":
			fmt.Println("呵呵哒")
			break
		default:
			fmt.Println("我也不知道你叫什么~")
		}
	}
}
