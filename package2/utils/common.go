package utils

import "fmt"

func init() {
	fmt.Println("common init")
}

func Add(a int, b int) int {
	return a + b
}