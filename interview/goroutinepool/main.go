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
	wp.Handle(func() {
		fmt.Println("run goroutine")
	})
}

type WorkerPool struct {
	Ready           []*WorkerChan
	WorkersCount    int
	MaxWorkersCount int

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

func (wp *WorkerPool) Handle(f func()) {
	worker := wp.getActiveWorker()
	worker.wc <- f
}

func (wp *WorkerPool) getActiveWorker() *WorkerChan {
	createWorker := false
	workerChan := new(WorkerChan)
	n := len(wp.Ready) - 1
	if n < 0 {
		// create a worker
		if wp.WorkersCount < wp.MaxWorkersCount {
			createWorker = true
			wp.WorkersCount++
		}
	} else {
		workerChan = wp.Ready[n]
		wp.Ready[n] = nil
		wp.Ready = wp.Ready[:n]
	}

	if workerChan == nil {
		if !createWorker {
			return nil
		}
		vch := wp.workerChanPool.Get()
		workerChan = vch.(*WorkerChan)
		// handle
		go func() {
			wp.WorkerFunc(workerChan)
			wp.workerChanPool.Put(workerChan)
		}()
	}

	return workerChan
}

func (wp *WorkerPool) WorkerFunc(wc *WorkerChan) {

}
