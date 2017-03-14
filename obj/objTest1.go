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

func main() {
	animal := Bird{10}
	animal.move()
	animal.move()
	animal.move()
	animal.move()
}
