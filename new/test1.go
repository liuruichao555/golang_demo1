package main

import "fmt"

func main() {
	b := new(bool)
	fmt.Println(*b)

	i := new(int)
	fmt.Println(*i)

	s := new(string)
	fmt.Println(*s)
}