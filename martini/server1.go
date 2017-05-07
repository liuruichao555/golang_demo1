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

	m.Get("/error", func() (int, string){
		return 404, "error"
	})

	m.RunOnAddr(":8080")
	//m.Run()
}