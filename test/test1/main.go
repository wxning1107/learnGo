package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func test1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover1: %v", err)
		}
	}()

	var mu sync.Mutex
	mu.Lock()
	mu.Unlock()
	mu.Unlock()
	time.Sleep(time.Second)
	err := recover()
	fmt.Printf("recover2: %v", err)
}

func test2() {
	m := map[int]string{}
	for i := 0; i < 10000; i++ {
		m[i] = strconv.Itoa(i)
	}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			//m[i] = strconv.Itoa(i + 1)
			fmt.Println(m[i])
		}(i)
	}
	time.Sleep(time.Second)
}

type Config struct {
	gin  func(g *[]string)
	name string
	defaultConfig
}
type defaultConfig struct {
	name string
}

func (d defaultConfig) get() {
	panic("implement me")
}

func (d defaultConfig) set() {
	panic("implement me")
}

type configInterface interface {
	get()
	set()
}

func main() {
	c := Config{}
	c.name = "abc"
	c.gin = func(g *[]string) {
		for i := 0; i < 10; i++ {
			*g = append(*g, strconv.Itoa(i))
			fmt.Println(g)
		}
	}

	time.Sleep(time.Second)
	g2 := []string{"a", "b"}
	c.gin(&g2)
	res := fmt.Sprintf("%.1f", 1.222)
	fmt.Println(res)

}
