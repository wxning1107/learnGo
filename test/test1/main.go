package main

import (
	"fmt"
	"time"
)

func main() {
	a := new(DescribeServiceResp)
	//a.Status.LoadBalancerIP = "a"
	fmt.Println(a.Status.LoadBalancerIP)
	fmt.Println(a.Status.Endpoints)

	//fmt.Println((*a.Status.Endpoints).formatTime())
	//a.Status.Endpoints.Format("2016")
	//t := new(Time)
	////t.Time = time.Time{}
	//fmt.Println(t.Time)
	//s := timeff(t.Time)
	//fmt.Println(s)
	testtt(a.Status.Endpoints)
	aa := ServiceEndpoints{}
	fmt.Println("aaaaaaaaaa")
	testtt(&aa)
	testtt(aa)
}

func testtt(e interface{}) {
	switch e.(type) {
	case ServiceEndpoints:
		fmt.Println("a")
	case *ServiceEndpoints:
		fmt.Println("b")
	}
}

type DescribeServiceResp struct {
	Status ServiceStatus `json:"status"`
}

type ServiceStatus struct {
	Endpoints      *ServiceEndpoints `json:"endpoints"`
	LoadBalancerIP string            `json:"loadBalancer_ip"`
	ClusterIP      string            `json:"cluster_ip"`
}

type ServiceEndpoints struct {
	time.Time
}

func (s *ServiceEndpoints) formatTime() string {
	//if s.IsZero() {
	//	return ""
	//}
	if s == nil {
		return "abc"
	}

	return s.Format("2006-01-02 15:04:05")
}

type Time struct {
	time.Time `protobuf:"-"`
}

func timeff(a interface{}) string {
	_, ok := a.(time.Time)
	if !ok {
		return "nnn"
	}
	return "NN"
}
