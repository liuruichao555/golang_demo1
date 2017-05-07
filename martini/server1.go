package main

import (
	"github.com/go-martini/martini"
	"github.com/ethereum/go-ethereum/log"
	"fmt"
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

	m.Get("/hello/**", func(params martini.Params) string {
		return "Hello " + params["_1"]
	})

	m.Get("/hello/(?P<name>[a-zA-Z]+)", func(params martini.Params) string {
		return fmt.Sprintf ("Hello %s", params["name"])
	})

	m.RunOnAddr(":8080")
	//m.Run()
}
