package main

import (
	"fmt"
	"time"
)

type User struct {
	username string
}

func (this *User) Close() {
	fmt.Println(this.username, "Closed!!!")
}

func main() {
	user := &User{"liuruichao"}
	defer user.Close()

	user2 := &User{"liuruichao2"}
	defer user2.Close()

	time.Sleep(10 * time.Second)
	fmt.Println("done")
}
