package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type User struct {
	username string  `json:"username"`
	password string  `json:"password"`
	age      int     `json:"age"`
}

func main() {
	var user = User{username: "liuruichao", password: "heheda", age: 20}

	// json编码
	body, err := json.Marshal(user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(body)

	// json解码
	str, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(str)
}
