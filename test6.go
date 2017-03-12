package main

import "fmt"

func main() {
	for {
		fmt.Print("Please input you simple name: ");
		var name string
		fmt.Scanln(&name)

		// exit
		if ("bye" == name) {
			break
		}

		if "lrc" == name {
			fmt.Println("刘瑞超")
		} else if "bzd" == name {
			fmt.Println("不知道")
		} else if "hhd" == name {
			fmt.Println("呵呵哒")
		} else {
			fmt.Println("我也不知道你叫什么")
		}
	}
}