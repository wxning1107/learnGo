package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type RWMutex struct {
	w           sync.Mutex
	mu          sync.Mutex
	readerCount int32
	readerWait  int32
}

const rwMutexMaxReaders = 1 << 30

func (rw *RWMutex) RLock() {
	//if rw.readerWait > 0 {
	//	rw.mu.Lock()
	//}
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		rw.w.Lock()
	}
}

func (rw *RWMutex) RUnlock() {
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		if r+1 == 0 || r+1 == -rwMutexMaxReaders {
			panic("sync: RUnlock of unlocked RWMutex")
		}

		if atomic.AddInt32(&rw.readerWait, -1) == 0 {

		}

		if rw.readerCount == 0 {
			rw.w.Unlock()
		}
	}
}

var sum int

func main() {
	var mu RWMutex
	sum = 0
	for i := 0; i < 1000; i++ {
		mu.RLock()
		go func() {
			sum++
		}()
		mu.RUnlock()
	}
	fmt.Println(sum)
}
