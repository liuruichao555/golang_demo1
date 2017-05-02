package main

import (
	"fmt"
	"time"
)

type Person struct {
	name       string
	age        int
	createTime int64
}

func main() {
	var _slice []int = make([]int, 5, 10)
	fmt.Println(_slice)

	var _slice1 []int = make([]int, 5)
	fmt.Println(_slice1)

	var _slice2 []int = []int{1, 2}
	fmt.Println(_slice2)

	person := Person{name: "liuruichao", age: 20, createTime: time.Now().Unix()}
	fmt.Println(person)

	var person2 map[string]Person = make(map[string]Person)
	person2["1"] = person

	fmt.Println(person2)

	fmt.Println(person2["1"])
}
