package main

import (
	"strings"
	"fmt"
)

func main() {
	var result string = "liuruichao"
	{
		tmp := "heheda"
		fmt.Println(strings.HasSuffix(result, "chao"))
		fmt.Println("tmp: ", tmp)
	}
	fmt.Println(result)
}