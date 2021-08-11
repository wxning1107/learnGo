package main

func Slice() {
	// 栈空间不足逃逸
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}

func main() {
	Slice()
}
