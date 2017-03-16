package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) (int, error) {
	return a + b, errors.New("计算错误")
}

func main() {
	result, err := add(1, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
