package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learnGoSource/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
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

	fmt.Println("start")
	p := PrintNum()
	for i := range p {
		fmt.Println(i)
	}
	fmt.Println("stop")

	fileSize, err := utils.GetFileSize("test.txt")
	fmt.Println(fileSize)
	fmt.Println(err)
	f, err = os.Open("test.txt")
	if err == nil {
		stat, _ := f.Stat()
		fmt.Println(stat.Size())
	}
	m := map[string]string{}
	err = json.Unmarshal([]byte(""), m)
	fmt.Println(err)

	s := utils.JoinString("abc", "def")
	fmt.Println(s)
	ParseTStringToTime("2020.08.08", "2020.02.09")
	TestReplace()
	fmt.Println(filepath.Join("home", ".kube", "config"))
}

func testLog() {
	log.Println("ccc")
}

func PrintNum() <-chan int {
	var wg sync.WaitGroup
	wg.Add(100)
	res := make(chan int, 100)
	ch := make(chan struct{}, 1)
	for i := 1; i <= 100; i++ {
		ch <- struct{}{}
		if i%2 == 0 {
			go func(i int) {
				res <- i
				<-ch
				wg.Done()
			}(i)
		} else {
			go func(i int) {
				res <- i
				<-ch
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
	close(res)
	close(ch)
	return res
}

func ParseTStringToTime(s1, s2 string) {
	t1, _ := time.Parse("2006-01-02", s1)
	t2, _ := time.Parse("2006-01-02", s1)
	t := t1.Sub(t2)
	fmt.Println(t.Hours())
	fmt.Println("```````````````````")
	fmt.Println(t1.Day())
}

func TestReplace() {
	s := `{"name":"{{name}}", "age":"{{age}}"}`
	replacer := strings.NewReplacer("{{name}}", "wxning")
	res := replacer.Replace(s)
	fmt.Println(res)
}
