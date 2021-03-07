package main

import (
	"bytes"
	"fmt"
)

func main() {
	sliceAppend()
	sliceAppendNew()
}

func sliceAppend() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')
	dir1 := path[:sepIndex]
	dir2 := path[sepIndex+1:]
	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1))
	fmt.Println("dir2 =>", string(dir2))
}

func sliceAppendNew() {
	s1 := make([]int, 8)
	s2 := s1[:4]
	s1 = append(s1, 1)
	s1[2] = 2
	fmt.Println(s1)
	fmt.Println(s2)
}
