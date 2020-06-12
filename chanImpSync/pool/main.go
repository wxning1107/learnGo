package main

import "fmt"

type Pool struct {
	buf   chan interface{}
	alloc func() interface{}
	clean func(interface{}) interface{}
}

func NewPool(size int, alloc func() interface{}, clean func(interface{}) interface{}) *Pool {
	return &Pool{
		buf:   make(chan interface{}, size),
		alloc: alloc,
		clean: clean,
	}
}

func (p *Pool) Get() interface{} {
	select {
	case i := <-p.buf:
		if p.clean != nil {
			return p.clean(i)
		}
		return i
	default:
		return p.alloc()

	}
}

func (p *Pool) Put(x interface{}) {
	select {
	case p.buf <- x:
	default:

	}
}

func main() {
	p := NewPool(1024,
		func() interface{} {
			return "no value"
		},
		func(i interface{}) interface{} {
			return "cleaned: " + i.(string)
		},
	)
	p.Put("aa")
	p.Put("bb")
	fmt.Println(p.Get())
	fmt.Println(p.Get())
	fmt.Println(p.Get())
}
