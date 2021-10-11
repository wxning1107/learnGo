package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// new worker pool
	wp := &WorkerPool{}
	wp.Start()

	// add func()
	wp.Go(func() {
		fmt.Println("run goroutine")
	})
}

type WorkerPool struct {
	Ready []*WorkerChan

	workerChanPool sync.Pool
}

type WorkerChan struct {
	wc          chan func()
	lastUseTime time.Time
}

var workerChanCap = func() int { return 1 }()

func (wp *WorkerPool) Start() {
	wp.workerChanPool.New = func() interface{} {
		return &WorkerChan{
			wc: make(chan func(), workerChanCap),
		}
	}

	go func() {
		for {
			// clean g that long time unused
			wp.Clean()
			time.Sleep(time.Second * 10)
		}
	}()
}

func (wp *WorkerPool) Clean() {
	for _, wc := range wp.Ready {
		if time.Now().Add(time.Second * 10).After(wc.lastUseTime) {
			// delete
		}
	}
}

func (wp *WorkerPool) Go(f func()) {
	worker := wp.getActiveWorker()
	worker.wc <- f
}

func (wp *WorkerPool) getActiveWorker() *WorkerChan {

}
