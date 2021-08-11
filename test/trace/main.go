package main

import (
	"os"
	"runtime"
	"runtime/trace"
)

// go run main.go 2> trace.out
// go tool trace trace.out
func main() {
	runtime.GOMAXPROCS(1)
	f, _ := os.Create("trace.out")

	_ = trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 100; i++ {
		_ = make([]int, 0, 40000)
	}

	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 100; i++ {
	//		_ = make([]int, 0, 10000)
	//	}
	//}()

}
