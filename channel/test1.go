package main

import (
	"fmt"
)

func main() {
	//var ch chan int
	//var ch1 chan <- int // ch1 只读
	//var ch2 <-chan int  // ch2 只写

	c := make(chan string)

	go func() {
		fmt.Println("呵呵哒")

		c <- "liuruichao"
	}()

	i := <- c
	fmt.Println(i)
}