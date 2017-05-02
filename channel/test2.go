package main

import (
	"fmt"
	"time"
)

func producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
		fmt.Println("send time: ", time.Now().Unix())
	}
}

func consumer(queue <- chan int) {
	for i := 0; i < 10; i++ {
		v := <- queue
		fmt.Printf("receive: %s, time: %s.\n", v, time.Now().Unix())
	}
}

func main() {
	queue := make(chan int, 1)
	go producer(queue)
	go consumer(queue)

	time.Sleep(1e9) //让Producer与Consumer完成
}
