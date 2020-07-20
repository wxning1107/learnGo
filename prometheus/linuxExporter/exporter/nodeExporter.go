package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"learnGoSource/prometheus/linuxExporter/scrapefunc"
	"sync"
)

var (
	reqCount int32
	hostName string
)

type LinuxExporter struct {
	portStatus *prometheus.Desc
}

func NewLinuxExporter(arg string) prometheus.Collector {
	return &LinuxExporter{
		prometheus.NewDesc(
			"clustermanager_oom_crashes_total",
			"Number of OOM crashes.",
			[]string{"host"},
			prometheus.Labels{"zone": arg},
		),
	}
}

type nodeStatsMetrics []struct {
	desc *prometheus.Desc
	eval func()
}

var wg sync.WaitGroup

func (n *LinuxExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.portStatus
}

func (n *LinuxExporter) Collect(ch chan<- prometheus.Metric) {
	go func() {
		for service, hostStatus := range scrapefunc.ScrapePortStatus() {
			for k, v := range hostStatus {
				ch <- prometheus.MustNewConstMetric(n.portStatus, prometheus.CounterValue, v, service, k)
			}
		}
	}()
}
