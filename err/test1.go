package main

import "fmt"

func add(a int, b int) int {
	if a == 0 || b == 0{
		panic(0)
	}
	return a + b
}

func main() {
	// 只有panic的异常才能被执行，相当于java里的finnaly，但是得在异常前声明
	defer func() {
		fmt.Println("捕获到异常")
	}()

	add(5, 3)
}