package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	logger "github.com/prometheus/common/log"
	LinuxExporter "learnGoSource/prometheus/cpuPortExporter/linuxExporter"
	"log"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8000", "Scrape port value")
	flag.Parse()

	exporter := LinuxExporter.NewLinuxExporter(port, "linux")
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(exporter)

	gatherers := prometheus.Gatherers{prometheus.DefaultGatherer, reg}
	handler := promhttp.HandlerFor(
		gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      logger.NewErrorLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(w, req)
	})

	log.Fatal(http.ListenAndServe(":1234", nil))
}
