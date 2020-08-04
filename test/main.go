package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	file, _ := os.Create("test.txt")
	defer file.Close()
	log.Printf("bbb")
	log.SetOutput(file)
	log.Printf("aaa")
	testLog()

	con, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("file content: %s\n", string(con))

	f, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("scanner: %s\n", scanner.Text())
	}

	var fl float64
	fl = 3.444
	fmt.Println(int(fl))
}

func testLog() {
	log.Println("ccc")
}
