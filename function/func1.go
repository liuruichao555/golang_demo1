package main

import "fmt"

func sort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if (arr[i] > arr[j]) {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func print(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d, ", arr[i])
	}
	fmt.Println()
}

func main() {
	arr := []int{5, 3, 1, 2, 10, -5}
	fmt.Println("========排序前=========")
	print(arr)
	fmt.Println("========排序后=========")
	sort(arr)
	print(arr)
}
