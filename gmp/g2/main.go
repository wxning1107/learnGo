package main

import (
	"fmt"
	"runtime"
	"time"
)

// go1.13输出：0 1 2 3 4 5 6 7 8 9
// go1.14输出：9 0 1 2 3 4 5 6 7 8
// 详情见：https://qcrao.com/2021/05/24/confusing-goroutine-running-orders/
func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Hour)
}
