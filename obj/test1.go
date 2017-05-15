package main

import "fmt"

func main() {
	type T1 struct {
		X int `json:"foo"`
	}

	type T2 struct {
		X int `json:"bar"`
	}

	var v1 T1
	var v2 T2

	v1 = T1{X: 1}
	v2 = T2{X: 2}

	v1 = T1(v2)

	fmt.Println(v1)
	fmt.Println(v2)
}
