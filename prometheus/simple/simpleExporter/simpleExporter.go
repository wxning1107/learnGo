package simpleExporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
)

type simpleExporter struct {
	zone         string
	OOMCountDesc *prometheus.Desc
	RAMUsageDesc *prometheus.Desc
}

func NewSimpleExporter(zone string) *simpleExporter {
	return &simpleExporter{
		zone:         zone,
		OOMCountDesc: prometheus.NewDesc("clustermanager_oom_crashes_total", "Number of OOM crashes.", []string{"host"}, prometheus.Labels{"zone": zone}),
		RAMUsageDesc: prometheus.NewDesc("clustermanager_ram_usage_bytes", "RAM usage as reported to the cluster manager.", []string{"host"}, prometheus.Labels{"zone": zone}),
	}
}

func ScrapeData() (oomCountByHost map[string]int, ramUsageByHost map[string]float64) {
	oomCountByHost = map[string]int{
		"foo.example.org": int(rand.Int31n(1000)),
		"bar.example.org": int(rand.Int31n(1000)),
	}
	ramUsageByHost = map[string]float64{
		"foo.example.org": rand.Float64() * 100,
		"bar.example.org": rand.Float64() * 100,
	}
	return
}

func (se *simpleExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- se.OOMCountDesc
	ch <- se.RAMUsageDesc
}

func (se *simpleExporter) Collect(ch chan<- prometheus.Metric) {
	oomCountByHost, ramUsageByHost := ScrapeData()
	for host, oomCount := range oomCountByHost {
		ch <- prometheus.MustNewConstMetric(se.OOMCountDesc, prometheus.CounterValue, float64(oomCount), host)
	}

	for host, ramUsage := range ramUsageByHost {
		ch <- prometheus.MustNewConstMetric(se.RAMUsageDesc, prometheus.CounterValue, ramUsage, host)
	}
}
