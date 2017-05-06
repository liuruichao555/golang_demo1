package main

import (
	"net/http"
	"log"
)

func WriteByte(body string) []byte {
	return []byte(body)
}

func handle(resp http.ResponseWriter, req *http.Request) {
	// 解析参数，默认不解析
	req.ParseForm()
	usernameArr := req.Form["username"]
	passwordArr := req.Form["password"]
	if len(usernameArr) <= 0 || len(passwordArr) <= 0 {
		resp.Write(WriteByte("参数错误"))
		return
	}
	username := usernameArr[0]
	password := passwordArr[0]
	if "liuruichao" == username && password == "liuruichao" {
		resp.Write(WriteByte("login success"))
		return
	}
	resp.Write(WriteByte("login fail"))
}

func main() {
	http.HandleFunc("/", handle)
	log.Println("server start ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}