package main

import "fmt"

type Animal interface {
	move()
}

type Bird struct {
	i int
}

func (r *Bird) move() {
	fmt.Printf("i: %d, %s.\n", r.i, "鸟类行走")
	r.i++
}

func (r *Bird) print() {
	fmt.Println("print: ", r)
}

func main() {
	animal := Bird{10}
	animal.move()
	animal.move()
	animal.move()
	animal.move()

	animal.print()
}
