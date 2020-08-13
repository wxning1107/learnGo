package main

import "fmt"

type Once struct {
	done chan struct{}
}

func NewOnce() Once {
	o := make(chan struct{}, 1)
	o <- struct{}{}

	return Once{done: o}
}

func (o *Once) Do(f func()) {
	if _, ok := <-o.done; !ok {
		return
	}

	f()

	close(o.done)
}

func main() {
	once := NewOnce()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		once.Do(func() {
			fmt.Println("Hello once")
		})
	}
}
