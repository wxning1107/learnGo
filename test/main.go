package main

import (
	"fmt"
)

type In interface {
	start() error
}

func main() {
	b := [2]int{6, 7}
	b[1] = 9
	fmt.Println(b)
	a := map[string][2]int{"2": [2]int{1, 2}}
	fmt.Println(a)
}
