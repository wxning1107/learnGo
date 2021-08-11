package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

// go tool pprof http://127.0.0.1:9999/debug/pprof/goroutine
func main() {
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func() {
			demo1()
			wg.Done()
		}()
	}
	wg.Wait()

	http.ListenAndServe("localhost:9999", nil)
}

func demo1() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Millisecond * 1200)
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
