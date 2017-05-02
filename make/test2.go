package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var _slice []int = make([]int, 5, 10)
	fmt.Println(_slice)

	var _slice1 []int = make([]int, 5)
	fmt.Println(_slice1)

	var _slice2 []int = []int{1, 2}
	fmt.Println(_slice2)
}