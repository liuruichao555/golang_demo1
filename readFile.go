package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func ReadAll(filePath string) ([]byte, error){
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func main() {
	buffer, err := ReadAll("test.dat")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(buffer))
}
