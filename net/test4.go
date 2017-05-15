package main

import (
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	url := "http://localhost:8080/test"

	reader := strings.NewReader("username=liuruichao&password=liuruichao")
	resp, err := http.Post(url, "application/x-www-form-urlencoded", reader)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	} else {
		fmt.Println("请求失败，请稍后重试")
	}
}