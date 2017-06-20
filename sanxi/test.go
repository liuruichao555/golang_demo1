package main

import (
	"fmt"
	"crypto/sha256"
	"io"
	"encoding/hex"
)

func main() {
	fmt.Println("start")
	hash := sha256.New()
	value := "liuruichao"
	io.WriteString(hash, value)
	signature := hash.Sum(nil)

	fmt.Println(hex.EncodeToString(signature))
}
