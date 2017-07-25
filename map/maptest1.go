package _map

import "fmt"

func main() {
	userMap := make(map[string]string)
	userMap["name"] = "liuruichao"
	userMap["age"] = "20"
	for k, v := range userMap {
		fmt.Printf("key: %s, value: %s.\n", k, v)
	}
}
