package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func main() {
	animal := Animal{name: "liuruichao", age: 20}
	fmt.Println(animal)
}