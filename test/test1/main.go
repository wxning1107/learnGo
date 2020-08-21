package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func test1() {
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

func test2() {
	m := map[int]string{}
	for i := 0; i < 10000; i++ {
		m[i] = strconv.Itoa(i)
	}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			//m[i] = strconv.Itoa(i + 1)
			fmt.Println(m[i])
		}(i)
	}
	time.Sleep(time.Second)
}
func main() {
	test2()
}
