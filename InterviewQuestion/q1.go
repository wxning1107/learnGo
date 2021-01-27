package main

import "fmt"

func main() {
	p := live()
	fmt.Println(p == nil)
}

type People interface {
	Show()
}

type Student struct{}

func (s *Student) Show() {}

func live() People {
	var s *Student
	return s
}
