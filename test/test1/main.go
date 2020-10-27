package main

import "fmt"

func main() {
	a := new(DescribeServiceResp)
	//a.Status.LoadBalancerIP = "a"
	fmt.Println(a.Status.LoadBalancerIP)
}

type DescribeServiceResp struct {
	Status ServiceStatus `json:"status"`
}

type ServiceStatus struct {
	Endpoints      []ServiceEndpoints `json:"endpoints"`
	LoadBalancerIP string             `json:"loadBalancer_ip"`
	ClusterIP      string             `json:"cluster_ip"`
}

type ServiceEndpoints struct {
	IP       string `json:"ip"`
	NodeName string `json:"node_name"`
}
