package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Int()
	c := make(chan int, 1)
	close(c)
	//c <- 2
	select {
	case v := <-c:
		fmt.Println(v)
	default:
		fmt.Println("aaa")
	}

}

type DescribeCronjobResp struct {
	Status CronjobStatus `json:"status"`
}

type CronjobStatus struct {
	Name             string `json:"name"`
	LastScheduleTime string `json:"last_schedule_time"`
}
