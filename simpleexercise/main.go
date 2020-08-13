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

func BubbleSort(s []int) {
	flag := false
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func MaxProfit(price []int) int {
	maxProfit := 0
	for i := 0; i < len(price)-1; i++ {
		if price[i+1]-price[i] > 0 {
			maxProfit += price[i+1] - price[i]
		}
	}
	return maxProfit
}

func MaxProfitOneTranc(p []int) int {
	MaxProfit := 0
	for i := 0; i < len(p); i++ {

	}

	return MaxProfit
}

func main() {
	printOddEven2()
	maxProfit := MaxProfit([]int{7, 1, 5, 3, 4, 6, 10})
	//maxProfit = MaxProfitOneTranc([]int{7, 1, 5, 3, 6, 4, 100})
	fmt.Println(maxProfit)
}
