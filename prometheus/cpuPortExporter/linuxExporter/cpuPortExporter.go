package LinuxExporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"learnGoSource/prometheus/cpuPortExporter/scrape"
	"sync"
)

const (
	namespace = "Linux"
	exporter  = "exporter"
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
			prometheus.BuildFQName(namespace, exporter, "port_status"),
			"Scrape port status: 2 is ESTABLISHED, 1 is LISTEN.",
			[]string{"port"},
			prometheus.Labels{"os": os},
		),
		CpuUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, exporter, "cpu_usage"),
			"Cpu usage.",
			[]string{"type"},
			prometheus.Labels{"os": os},
		),
	}
}

func (me *LinuxExporter) Describe(ch chan<- *prometheus.Desc) {
	desc := make(chan prometheus.Metric)
	doneChan := make(chan struct{})
	go func() {
		for metric := range desc {
			ch <- metric.Desc()
		}
		doneChan <- struct{}{}
		close(doneChan)
	}()

	me.Collect(desc)
	close(desc)
	<-doneChan
}

func (me *LinuxExporter) Collect(ch chan<- prometheus.Metric) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for port, hostStatus := range scrape.ScrapePortStatus(me.Port) {
			ch <- prometheus.MustNewConstMetric(me.PortStatus, prometheus.CounterValue, hostStatus, port)
		}
		wg.Done()
	}()

	go func() {
		for CpuType, usage := range scrape.ScrapeCpuUsage() {
			ch <- prometheus.MustNewConstMetric(me.CpuUsage, prometheus.CounterValue, usage, CpuType)
		}
		wg.Done()
	}()
	wg.Wait()

	// For test
	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//	for port, hostStatus := range scrape.ScrapeMacPort(me.Port, ch) {
	//		ch <- prometheus.MustNewConstMetric(me.PortStatus, prometheus.CounterValue, hostStatus, port)
	//	}
	//	wg.Done()
	//}()
	//
	//go func() {
	//	for cpu, usage := range scrape.ScrapeMacCpu() {
	//		ch <- prometheus.MustNewConstMetric(me.CpuUsage, prometheus.CounterValue, usage, cpu)
	//	}
	//	wg.Done()
	//}()

	wg.Wait()
}
