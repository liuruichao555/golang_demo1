package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
}

func main() {
	jsonStr := `{"username": "liuruichao", "password": "heheda", "age": 20}`
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("username: %s, password: %s, age: %d.\n", user.Username, user.Password, user.Age)

	body, err := json.Marshal(user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))
}
