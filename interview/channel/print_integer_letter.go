package main

import "fmt"

// 开两个goroutine交替打印数字和字母
func main() {
	strSource := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterChan := make(chan struct{})
	integerChan := make(chan struct{})
	done := make(chan struct{})
	n := 0

	go func() {
		integerChan <- struct{}{}
	}()
	go func() {
		i := 1
		for {
			select {
			case <-integerChan:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				if n < 25 {
					letterChan <- struct{}{}
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-letterChan:
				fmt.Print(strSource[n : n+1])
				n++
				fmt.Print(strSource[n : n+1])
				n++
				integerChan <- struct{}{}
				if n == 26 {
					close(done)
				}
			}
		}
	}()

	select {
	case <-done:
		return
	}
}
