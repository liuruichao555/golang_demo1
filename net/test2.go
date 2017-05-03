package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	client := &http.Client{}
	url := "http://www.baidu.com"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		html, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(html))
	} else {
		fmt.Println("访问失败")
	}
}