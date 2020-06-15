package main

import (
	"fmt"
)

type generation struct {
	// A barrier for waiters to wait on.
	// This will never be used for sending, only receive and close.
	wait chan struct{}
	// The counter for remaining jobs to wait for.
	n int
}

func newGeneration() generation {
	return generation{wait: make(chan struct{})}
}
func (g generation) end() {
	// The end of a generation is signalled by closing its channel.
	close(g.wait)
}

// Here we use a channel to protect the current generation.
// This is basically a mutex for the state of the WaitGroup.
type WaitGroup chan generation

func NewWaitGroup() WaitGroup {
	wg := make(WaitGroup, 1)
	g := newGeneration()
	// On a new waitgroup waits should just return, so
	// it behaves exactly as after a terminated generation.
	g.end()
	wg <- g
	return wg
}

func (wg WaitGroup) Add(delta int) {
	// Acquire the current generation.
	g := <-wg
	if g.n == 0 {
		// We were at 0, create the next generation.
		g = newGeneration()
	}
	g.n += delta
	if g.n < 0 {
		// This is the same behavior of the stdlib.
		panic("negative WaitGroup count")
	}
	if g.n == 0 {
		// We reached zero, signal waiters to return from Wait.
		g.end()
	}
	// Release the current generation.
	wg <- g
}

func (wg WaitGroup) Done() { wg.Add(-1) }

func (wg WaitGroup) Wait() {
	// Acquire the current generation.
	g := <-wg
	// Save a reference to the current waiting chan.
	wait := g.wait
	// Release the current generation.
	wg <- g
	// Wait for the chan to be closed.
	<-wait
}

func main() {
	wg := NewWaitGroup()
	v := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			v++
			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(time.se)
	fmt.Println(v)
}
