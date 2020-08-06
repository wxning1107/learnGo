package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printOddEven() {
	var wg sync.WaitGroup
	wg.Add(100)
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				fmt.Println(i)
				wg.Done()
				runtime.Gosched()
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 1 {
				fmt.Println(i)
				wg.Done()
				runtime.Gosched()
			}
		}
	}()
	wg.Wait()
}

func printOddEven2() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 0 {
				ch <- i
				wg.Done()
				fmt.Println(i)
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			if i%2 == 1 {
				<-ch
				wg.Done()
				fmt.Println(i)
			}
		}
	}()
	wg.Wait()
}

func main() {
	printOddEven2()
}
