package main

import "fmt"

type Person struct {
	name string
	// age int
}

type Dog struct {
	name string
}

func print(person Person, dog Dog) {
	fmt.Println("强制转换前")
	fmt.Println("\t", person)
	fmt.Println("\t", dog)
}

func main() {

	person := Person{name: "heheda"}
	dog := Dog{name: "xiaobai"}

	print(person, dog)

	dog = Dog(person)

	print(person, dog)
}
