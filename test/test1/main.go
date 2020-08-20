package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover1: %v", err)
		}
	}()

	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
	mu.Unlock()
	time.Sleep(time.Second)
	err := recover()
	fmt.Printf("recover2: %v", err)

}
