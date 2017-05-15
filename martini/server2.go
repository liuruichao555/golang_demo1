package main

import (
	"github.com/go-martini/martini"
	"fmt"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello World"
	})

	m.Post("/test", func(body string) string {
		fmt.Println(body)
		return "success"
		/*username := params["username"]
		password := params["password"]*/
		/*fmt.Println("username: ", username, ", password: ", password, ".")
		if username == "" || password == "" {
			return "username or password is empty!"
		}

		if username == "liuruichao" && password == "liuruichao" {
			return "login success"
		}
		return "login fail"*/
	})

	m.RunOnAddr(":8080")
}