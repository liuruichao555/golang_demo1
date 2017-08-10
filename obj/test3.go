package main

import "fmt"

type Person2 struct {
	name string
	age  int
}

func main() {
	person := Person2{name: "liuruichao", age: 20}
	fmt.Println(person)
}
