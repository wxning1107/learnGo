package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"
)

/**
go tool pprof http://localhost:1234/debug/pprof/profile?seconds=60
go tool pprof http://localhost:1234/debug/pprof/heap
*/
func main() {
	go func() {
		for {
			log.Println(add("www.wxning.com"))
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:1234", nil)
}

var datas []string

func add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}
