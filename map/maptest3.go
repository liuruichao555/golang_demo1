package main

import (
	"github.com/spf13/viper"
	"fmt"
)

func main() {
	viper.SetConfigName("core")
	viper.AddConfigPath("/Users/liuruichao/develop/opensource/golang/gopath/src/github.com/hyperledger/fabric/sampleconfig")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	for _, key := range viper.AllKeys() {
		fmt.Printf("value: %s, value: %v.\n", key, viper.GetString(key))
	}
}
