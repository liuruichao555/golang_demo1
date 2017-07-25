package main

import "fmt"

type User struct {
	name string
    age  int
}

func main() {
	userMap := make(map[string]User)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("lrc_%d", i)
		fmt.Println("key: ", key)
		userMap[key] = User{name: "liuruichao", age: 20 + i}
	}
	fmt.Printf("userMape size: %d.\n", len(userMap))

	for key, value := range userMap {
		fmt.Printf("key: %s, value: %v.\n", key, value)
	}
}