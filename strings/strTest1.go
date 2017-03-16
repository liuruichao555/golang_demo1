package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "liuruichao, liuruichao2, liuruichao3"
	fmt.Println(name[1: 2]) // 截取字符串

	strs := strings.Split(name, ",")
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}

	fmt.Println(strings.HasPrefix(name, "liuruichao"))
	fmt.Println(strings.HasSuffix(name, "liuruichao"))

	fmt.Println(strings.Contains(name, "liuruichao"))

	fmt.Println(name + ", heheda")
}
