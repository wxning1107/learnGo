package scrape

import (
	"github.com/prometheus/client_golang/prometheus"
	"learnGoSource/prometheus/cpuPortExporter/common"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func ScrapePortStatus(port string) map[string]float64 {
	res := make(map[string]float64)
	shell := `netstat -antup |grep ` + port
	scriptRes, err := common.ExeScript(shell)
	if err != nil {
		log.Printf("Error occured when scrape prot status: %v\n", err)
	}
	for _, portStatus := range strings.Split(strings.Trim(scriptRes, "\n"), " ") {
		switch portStatus {
		case "ESTABLISHED":
			res[port] = float64(2)
		case "LISTEN":
			res[port] = float64(1)
		default:
			res[port] = float64(0)
		}
	}

	return res
}

func ScrapeCpuUsage() map[string]float64 {
	res := make(map[string]float64)
	shell := `top`
	scriptRes, err := common.ExeScript(shell)
	if err != nil {
		log.Printf("Error occured when scrape cpu usage: %v\n", err)
	}
	part := strings.Split(strings.Trim(scriptRes, "\n"), "\n")[2]
	for i, usage := range part {
		if string(usage) == "us" {
			value, _ := strconv.ParseFloat(string(part[i-1]), 64)
			res["user_space"] = value
		}
		if string(usage) == "sy" {
			value, _ := strconv.ParseFloat(string(part[i-1]), 64)
			res["sysctl"] = value
		}
		if string(usage) == "id" {
			value, _ := strconv.ParseFloat(string(part[i-1]), 64)
			res["cpu_usage"] = float64(1) - value
		}
	}
	return res
}

func ScrapeMacPort(port string, ch chan<- prometheus.Metric) map[string]float64 {
	res := make(map[string]float64)
	res[port] = float64(1)
	return res
}

func ScrapeMacCpu() map[string]float64 {
	res := make(map[string]float64)
	res["user_space"] = rand.Float64() * 100
	res["sysctl"] = rand.Float64() * 100
	res["cpu_usage"] = rand.Float64() * 100
	return res
}
