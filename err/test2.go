package main

import "fmt"

func getNum() int {
	defer func() {
		fmt.Println("getNum()")
	}()
	panic(0)
	//a := 5 / 0 // 这个是不走defer
	return 0
}

func main() {
	a := getNum()
	fmt.Println(a)
}