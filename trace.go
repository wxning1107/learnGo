package main

import (
	"errors"
	"os"
	"runtime/trace"
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

//func main() {
//	//t := test{a:b{aa:1}}
//	//err := t.a.start()
//	//fmt.Println(err)
//	//tt := b{aa: 1}
//	//tt.start()
//	//fmt.Println(tt.aa)
//	//rand.Seed(time.Now().Unix())
//	//fmt.Println(rand.Intn(100))
//	//fmt.Println(rand.Intn(100))
//	runtime.GOMAXPROCS(1)
//	go func() {
//		fmt.Println(1)
//	}()
//	go func() {
//		fmt.Println(2)
//	}()
//	go func() {
//		fmt.Println(3)
//	}()
//	go func() {
//		time.Sleep(time.Second)
//	}()
//	//time.Sleep(time.Second)
//
//}

//func main() {
//	trace.Start(os.Stderr)
//	defer trace.Stop()
//	ch := make(chan string)
//	go func() {
//		ch <- "EDDYCJY"
//	}()
//	<-ch
//}

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	ch := make(chan string)
	go func() {
		ch <- "EDDYCJY"
	}()
	<-ch
}
