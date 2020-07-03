package main

import (
	"fmt"
	"time"
)

type In interface {
	start() error
}

func main() {
	c := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Millisecond)
}
