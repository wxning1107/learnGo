package main

import (
	"fmt"
	"time"
)

func main() {
	demo1()
	demo2()
	demo3()
}

// panic 只会触发当前 Goroutine 的延迟函数调用
// defer 关键字对应的 runtime.deferproc 会将延迟调用函数与调用方所在 Goroutine 进行关联。所以当程序发生崩溃时只会调用当前 Goroutine 的延迟调用函数
func demo1() {
	defer println("in main")
	go func() {
		defer println("in goroutine")
		panic("panic occurred")
	}()

	time.Sleep(1 * time.Second)
}

// recover 只有在发生 panic 之后调用才会生效。然而在上面的控制流中，recover 是在 panic 之前调用的，并不满足生效的条件，所以我们需要在 defer 中使用 recover 关键字
func demo2() {
	defer fmt.Println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unknown err")
}

// 多次调用 panic 也不会影响 defer 函数的正常执行，所以使用 defer 进行收尾工作一般来说都是安全的
func demo3() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
