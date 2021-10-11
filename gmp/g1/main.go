package main

import (
	"fmt"
	"runtime"
)

// 执行结果是 9, 0, 1, 2, 3, 4, 5, 6, 7, 8
// 原因是9是最后放到runnext中，其他都在p本地队列中，runnext会优先打印。详情见：https://qcrao.com/2021/05/24/confusing-goroutine-running-orders/
func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	var ch = make(chan int)
	<-ch
}
