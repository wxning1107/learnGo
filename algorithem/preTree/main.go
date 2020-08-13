package main

import (
	"fmt"
)

type Node struct {
	next   map[string]*Node
	isWord bool
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(word string) {
	cur := t.root
	for _, w := range []rune(word) {
		if cur.next[string(w)] == nil {
			cur.next[string(w)] = &Node{next: make(map[string]*Node)}
		}
		cur = cur.next[string(w)]
	}
	cur.isWord = true
}

func (t *Trie) Search(word string) bool {
	cur := t.root
	for _, w := range []rune(word) {
		if cur.next[string(w)] == nil {
			return false
		}
		cur = cur.next[string(w)]
	}
	return cur.isWord
}

func main() {
	t := Trie{
		root: &Node{next: make(map[string]*Node)},
	}
	t.Insert("wxning")
	isWord := t.Search("wxning")
	fmt.Println(isWord)
}
