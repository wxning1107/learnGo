package main

import "fmt"

func main() {
	head := &node{1, nil}
	head.next = &node{2, nil}
	fmt.Println("before main", head)
	modify(head)
	fmt.Println("after main", head)
}

func modify(head *node) {
	fmt.Println("before modify", head)
	*head = node{}
	fmt.Println("after modify", head)
}

type node struct {
	data int
	next *node
}
