package main

import "sync"

var counter int

func main() {
	sync.Mutex{}
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}

	wg.Wait()
	println(counter)
}
