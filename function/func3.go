package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a和b不能小于0")
	}
	return a + b, nil
}

func main() {
	//result, err := add(1, 3) // error
	result, err := add(-3, 1)  // success
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
