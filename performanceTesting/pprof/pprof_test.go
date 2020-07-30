package main

import "testing"

const url = "www.wxning.com"

func TestAdd(t *testing.T) {
	s := add(url)
	if s == "" {
		t.Errorf("Testing error")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(url)
	}
}
