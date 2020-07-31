package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type config struct {
	c1 struct {
		c1Config1 string
		c1Config2 string
	}
	c2 struct {
		c2Config1 string
		c2Config2 string
	}
}

func main() {
	var cfg config
	var c = config{
		c1: struct {
			c1Config1 string
			c1Config2 string
		}{c1Config1: "111", c1Config2: "222"},
		c2: struct {
			c2Config1 string
			c2Config2 string
		}{c2Config1: "333", c2Config2: "444"},
	}

	res, _ := json.Marshal(c)
	_ = json.Unmarshal(res, &cfg)
	fmt.Printf("%v\n", cfg)

	var buf bytes.Buffer
	log.Printf("bbb")
	file, _ := os.Open("file")
	log.SetOutput(file)
	log.Printf("aaa")
	testLog()
	fmt.Println(buf.String())

}

func testLog() {
	log.Println("ccc")
}
