package common

import (
	"github.com/prometheus/client_golang/prometheus"
	prometheusClient "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"strings"
)

func InitExporter(exporter prometheus.Collector) prometheus.Gatherers {
	var parser expfmt.TextParser
	var parserText = func() ([]*prometheusClient.MetricFamily, error) {
		parsed, err := parser.TextToMetricFamilies(strings.NewReader(""))
		if err != nil {
			return nil, err
		}
		var result []*prometheusClient.MetricFamily
		for _, mf := range parsed {
			result = append(result, mf)
		}
		return result, nil
	}

	foo := exporter
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(foo)
	newGatherers := prometheus.Gatherers{
		prometheus.GathererFunc(parserText),
		reg,
	}

	return newGatherers
}
