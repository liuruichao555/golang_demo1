package main

import "fmt"

type Animal interface {
	print()
}

type Dog struct {
	name string
	age int
}

func (dog *Dog) print() {
	fmt.Printf("name: %s, age: %d,", dog.name, dog.age)
}

func main() {
	dog := Dog{name: "xiaobao", age: 2}
	dog.print()
}