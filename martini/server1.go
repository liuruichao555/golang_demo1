package main

import (
	"github.com/go-martini/martini"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	m := martini.Classic()

	m.Get("/", func() string {
		log.Info("hello")
		return "Hello World!"
	})

	m.Get("/hello/:name", func(params martini.Params) string {
		return "welcome => " + params["name"]
	})

	m.Get("/error", func() (int, string) {
		return 404, "error"
	})

	m.NotFound(func() string {
		return "404"
	})

	m.RunOnAddr(":8080")
	//m.Run()
}
