package main

import (
	"github.com/go-martini/martini"
	"fmt"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello World"
	})

	m.Post("/test", func(req *http.Request) string {
		username := req.FormValue("username")
		password := req.FormValue("password")
		fmt.Println("username: ", username, ", password: ", password, ".")
		if username == "" || password == "" {
			return "username or password is empty!"
		}

		if username == "liuruichao" && password == "liuruichao" {
			return "login success"
		}
		return "login fail"
	})

	m.RunOnAddr(":8080")
}