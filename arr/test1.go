package main

import "fmt"

func main() {
	arr := [...] int {1, 2, 3, 4, 5}
	fmt.Printf("arr: %v\n", arr)

	// 切片类型
	s := arr[:]
	s2 := append(s, 10, 22, 33, 44)
	fmt.Printf("s: %v\n", s2)

	s = append(s, 100, 200, 300)
	fmt.Printf("s: %v\n", s)
}
