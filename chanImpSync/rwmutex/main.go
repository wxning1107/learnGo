package main

type RWMutex struct {
	writeChan chan struct{}
	readers   chan int
}

func NewRWMutex() RWMutex {
	return RWMutex{
		writeChan: make(chan struct{}, 1),
		readers:   make(chan int, 1),
	}
}

func (m RWMutex) Lock() {
	m.writeChan <- struct{}{}
}

func (m RWMutex) Unlock() {
	<-m.writeChan
}

func (m RWMutex) RLock() {
	var rs int

	select {
	case m.writeChan <- struct{}{}:
	case rs = <-m.readers:
	}

	rs++
	m.readers <- rs
}

func (m RWMutex) RUnlock() {
	rs := <-m.readers
	rs--

	if rs == 0 {
		<-m.writeChan
		return
	}

	m.readers <- rs
}

func main() {
	mu := NewRWMutex()
	//mu.Lock()
	mu.RLock()
	mu.RLock()
	mu.RLock()
	mu.RUnlock()
	mu.RUnlock()
	mu.RUnlock()
	//mu.Unlock()
}
