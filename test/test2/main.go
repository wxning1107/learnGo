package main

import (
	"bufio"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var num int64
var url = "user-system.stg.svc.qt-k8s-hz.com/v1/users/6967958489d6618f2b82a451aa76b495"

func main() {
	//TestAPPFrameworkQPS()
	//
	//fmt.Println(num)
	file, err := os.OpenFile("result.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	//defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("vvv")
	writer.Flush()
}

func TestAPPFrameworkQPS() {
	go func() {
		for {
			randMilliTime := rand.Intn(1000)
			time.Sleep(time.Millisecond * time.Duration(randMilliTime))

			resp, err := http.Get("http://www.baidu.com")
			if err != nil {
				log.Printf("get err: %v\n", err)
			}
			if resp.StatusCode != http.StatusOK {
				atomic.AddInt64(&num, 1)
				log.Printf("get wrong statuscode: %v, num: %v\n", resp.StatusCode, num)
			}
			resp.Body.Close()

			time.Sleep(time.Millisecond * time.Duration(1000-randMilliTime))
		}
	}()
}
