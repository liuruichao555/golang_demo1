package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	url := "http://python.51universe.com/test.php"

	reader := strings.NewReader("name=liuruichao&age=20")
	resp, err := http.Post(url, "application/x-www-form-urlencoded", reader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
