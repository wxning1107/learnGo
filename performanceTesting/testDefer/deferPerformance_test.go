package testDefer

import "testing"

/**
go test -bench=. -benchmem -run=none
go tool compile -S main.go

*/
func BenchmarkDoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoDefer("wxning", "www.wxning.com")
	}
}

func BenchmarkDoNotDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoNotDefer("wxning", "www.wxning.com")
	}
}

func DoDefer(key, value string) {
	defer func(key, value string) {
		_ = key + value
	}(key, value)
}

func DoNotDefer(key, value string) {
	_ = key + value
}
