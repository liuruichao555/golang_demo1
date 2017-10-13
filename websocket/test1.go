package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"log"
	"fmt"
)

func myHandle(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", msg[:n])

	sendMsg := "[" + string(msg[:n]) + "]"
	m, err := ws.Write([]byte(sendMsg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", msg[:m])
}

func main() {
	http.Handle("/echo", websocket.Handler(myHandle))
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
