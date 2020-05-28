package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type In interface {
	start() error
}
type test struct {
	a In
}
type b struct {
	aa int
}

func (b b) start() error {
	b.aa = 2
	return errors.New("bbb")
}
func main() {
	//t := test{a:b{aa:1}}
	//err := t.a.start()
	//fmt.Println(err)
	tt := b{aa: 1}
	tt.start()
	fmt.Println(tt.aa)
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
