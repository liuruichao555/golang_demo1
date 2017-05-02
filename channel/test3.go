package main

import "fmt"

func sum(values[] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}

	// 将计算结果发送到channel中
	resultChan <- sum
}

func main() {
	resultChan := make(chan int, 3)
	go sum([]int{1, 1, 1}, resultChan)
	go sum([]int{2, 2, 2}, resultChan)
	go sum([]int{3, 3, 3}, resultChan)
	sum1, sum2, sum3 := <- resultChan, <- resultChan, <-resultChan
	fmt.Println("result: ", sum1, sum2, sum3)
}
