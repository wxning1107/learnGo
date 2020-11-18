package main

import "fmt"

func main() {
	a := new(CronjobStatus)
	fmt.Println(string(nil))
}

type DescribeCronjobResp struct {
	Status CronjobStatus `json:"status"`
}

type CronjobStatus struct {
	Name             *string `json:"name"`
	LastScheduleTime string  `json:"last_schedule_time"`
}
