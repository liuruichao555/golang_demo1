package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	// 加锁
	mutex.Lock()

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Println("Lock the lock.", i)
			mutex.Lock()
			fmt.Println("The lock is locked", i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("unlock the lock")
	// 释放锁
	mutex.Unlock()

	fmt.Println("The lock is unlocked.")
	time.Sleep(time.Second)
}