package main

import (
	"fmt"
	"time"
)

func producer2(queue chan string) {
	for i := 0; i < 100; i++ {
		queue <- fmt.Sprintf("liuruichao_%d", i)
		fmt.Printf("send %d to queue.\n", i)
	}
}

func consumer2(queue chan string) {
	for {
		str := <-queue
		fmt.Println("consumer receive: " + str)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	queue := make(chan string, 100)

	go producer2(queue)
	go consumer2(queue)

	time.Sleep(300 * time.Second) //让Producer与Consumer完成
}
