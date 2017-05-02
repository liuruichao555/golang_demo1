package main

import "fmt"

func main() {
	var _map map[string]string = make(map[string]string)

	for i := 0; i < 10; i++ {
		_map["name_" + string(i)] = "liuruichao" + string(i)
	}

	fmt.Println(len(_map))

	for key, value := range _map{
		fmt.Printf("key: %s, value: %s.\n", key, value)
	}
}