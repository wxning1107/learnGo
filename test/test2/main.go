package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x int
	runtime.GOMAXPROCS(2)
	for i := 0; i < 2; i++ {
		go func() {
			x++
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(x)

}
