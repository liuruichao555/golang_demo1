package main

import (
	"time"
	"fmt"
)

func main() {
	timeout := make(chan bool, 1)

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		time.Sleep(1e9) // one second
		timeout <- true
	}()

	for i := 1; i < 5; i++ {
		select {
		case t1 := <-ch:
			fmt.Println("ch pop one element ", t1)
		case <-timeout:
			fmt.Println("timeout!")
		}
	}

	time.Sleep(5 * time.Second)
}
