package main

import "fmt"

func getNum() int {
	defer func() {
		fmt.Println("defer")
	}()
	return 0
}

func main() {
	a := getNum()
	fmt.Println(a)
}