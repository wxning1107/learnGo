package main

import (
	"io"
	"os"
	"sync"
	"testing"
)

func mockReadFile(b []byte) {
	f, _ := os.Open("water")
	for {
		n, err := io.ReadFull(f, b)
		if n == 0 || err == io.EOF {
			break
		}
	}
}

func BenchmarkNewBytePoolCap(b *testing.B) {
	bp := NewBytePoolCap(500, 1024, 1024)
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(bp *BytePoolCap) {
			buf := bp.Get()
			defer bp.Put(buf)
			mockReadFile(buf)
			wg.Done()
		}(bp)
	}
	wg.Wait()
}
