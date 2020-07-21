package linuxExporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"learnGoSource/prometheus/cpuPortExporter/scrape"
)

type LinuxExporter struct {
	OperationSystem string
	Port            string
	PortStatus      *prometheus.Desc
	CpuUsage        *prometheus.Desc
}

func NewLinuxExporter(port string, os string) prometheus.Collector {
	return &LinuxExporter{
		OperationSystem: os,
		Port:            port,
		PortStatus: prometheus.NewDesc(
			"port_status",
			"Scrape port status: 2 is ESTABLISHED, 1 is LISTEN.",
			[]string{"port"},
			prometheus.Labels{"os": os},
		),
		CpuUsage: prometheus.NewDesc(
			"cpu_usage",
			"Cpu usage",
			[]string{"type"},
			prometheus.Labels{"os": os},
		),
	}
}

func (me *LinuxExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- me.PortStatus
	ch <- me.CpuUsage
}

func (me *LinuxExporter) Collect(ch chan<- prometheus.Metric) {
	//for port, hostStatus := range scrape.ScrapePortStatus(me.Port) {
	//	ch <- prometheus.MustNewConstMetric(me.PortStatus, prometheus.CounterValue, hostStatus, port)
	//}
	//
	//for CpuType, usage := range scrape.ScrapeCpuUsage() {
	//	ch <- prometheus.MustNewConstMetric(me.CpuUsage, prometheus.CounterValue, usage, CpuType)
	//}
	for port, hostStatus := range scrape.ScrapeMacPort(me.Port) {
		ch <- prometheus.MustNewConstMetric(me.PortStatus, prometheus.CounterValue, hostStatus, port)
	}
	for cpu, usage := range scrape.ScrapeMacCpu() {
		ch <- prometheus.MustNewConstMetric(me.CpuUsage, prometheus.CounterValue, usage, cpu)
	}
}
