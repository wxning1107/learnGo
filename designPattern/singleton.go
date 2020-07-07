package main

import (
	"fmt"
	"sync"
)

type WebConfig struct {
	Port int
}

var wc *WebConfig
var mu sync.Mutex

func GetConfig() *WebConfig {
	mu.Lock()
	defer mu.Unlock()
	if wc == nil {
		wc = &WebConfig{Port: 8080}
	}
	return wc
}

var once sync.Once

func InitConfig() *WebConfig {
	once.Do(func() {
		wc = &WebConfig{Port: 8000}
	})
	return wc
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()
	fmt.Println(c1 == c2)

	wc = nil

	c3 := InitConfig()
	c4 := InitConfig()
	fmt.Println(c3 == c4)
}
