package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var ret []string
	ret = append(ret, "{\"name\":\"liuruichao\"}")
	ret = append(ret, "{\"name\":\"liuruichao2\"}")
	jsonStr, err := json.Marshal(ret)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(jsonStr)
}
