package main

import (
	"fmt"
	"time"
)

type Mutex struct {
	Semaphore chan struct{}
}

func NewMutex(size int) Mutex {
	return Mutex{make(chan struct{}, size)}
}

func (m Mutex) Lock() {
	m.Semaphore <- struct{}{}
}

//func (m Mutex) TryLock() bool {
//	select {
//	case m.Semaphore <- struct {}{}:
//		return true
//	default:
//		return false
//	}
//}

func (m Mutex) Unlock() {
	<-m.Semaphore
}

func main() {
	s := NewMutex(1)
	count := 0
	for i := 0; i < 10000; i++ {
		go func(i int) {
			s.Lock()
			count++
			s.Unlock()
		}(i)
	}
	time.Sleep(time.Millisecond * 5)
	fmt.Println(count)
}
