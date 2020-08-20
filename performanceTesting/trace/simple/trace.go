package main

import (
	"os"
	"runtime/trace"
)

/**
go run trace.go 2> trace.out
go tool trace trace.out
*/
func main() {
	_ = trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "WXNING"
	}()
	<-ch
}
