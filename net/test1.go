package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(html))
}