package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		return "Hello World!"
	})

	m.Get("/hello/:name", func(params martini.Params) string {
		return "welcome => " + params["name"]
	})

	m.RunOnAddr(":8080")
	//m.Run()
}