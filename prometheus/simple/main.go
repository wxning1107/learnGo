package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	logger "github.com/prometheus/common/log"
	"learnGoSource/prometheus/simple/simpleExporter"
	"log"
	"net/http"
)

func init() {
}

func main() {
	workerDB := simpleExporter.NewSimpleExporter("db")
	workerCA := simpleExporter.NewSimpleExporter("ca")
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(workerDB)
	reg.MustRegister(workerCA)

	gatherers := prometheus.Gatherers{prometheus.DefaultGatherer, reg}
	handler := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{ErrorLog: logger.NewErrorLogger(), ErrorHandling: promhttp.ContinueOnError})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(w, req)
	})
	log.Fatal(http.ListenAndServe(":1234", nil))
	//http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(":1234", nil))
}
