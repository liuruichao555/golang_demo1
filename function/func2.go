package main

import "fmt"

func add(num1 int, num2 int) (int, string) {
	return num1 + num2, "success"
}

func main() {
	a := 10
	b := 5
	fmt.Println(add(a, b))
}
