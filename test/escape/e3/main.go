package main

import "errors"

// 动态类型逃逸
func main() {
	//s := "Escape"
	//fmt.Println(s)
	var f = make(chan error, 1)
	go func() {
		f <- errors.New("a")
	}()
	select {
	case <-f:

	}
}

// 内联
