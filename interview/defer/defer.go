package main

import "fmt"

func main() {
	demo1()
}

func demo1() {
	var whatever [3]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}
