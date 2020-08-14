package main

import "testing"

func convertIntToInterface(val int) []interface{} {
	s := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		s[i] = val
	}
	return s
}

func BenchmarkConvertSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := convertIntToInterface(12)
		_ = res
	}
}

// go test -bench . -benchmem ./...
func BenchmarkConvertBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := convertIntToInterface(256)
		_ = res
	}
}
