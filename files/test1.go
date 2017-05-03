package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	// read file
	data, err := ioutil.ReadFile("/Users/liuruichao/shell/test/test4.sh")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(data))

	// write file
	err = ioutil.WriteFile("/Users/liuruichao/tmp/test.dat", []byte("liuruichao\n"), 0644)
	if err != nil {
		panic(err)
	}
}