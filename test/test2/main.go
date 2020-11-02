package main

import "fmt"

func main() {
	a := new(DescribeCronjobResp)
	fmt.Println(a.Status.Name)
}

type DescribeCronjobResp struct {
	Status CronjobStatus `json:"status"`
}

type CronjobStatus struct {
	Name             string `json:"name"`
	LastScheduleTime string `json:"last_schedule_time"`
}
