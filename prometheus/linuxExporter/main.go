package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"learnGoSource/prometheus/linuxExporter/exporter"
	"log"
	"net/http"
)

//func main() {
//	prometheus.MustRegister()
//	http.Handle("/metrics", promhttp.Handler())
//	log.Fatal(http.ListenAndServe(":1234", nil))
//}

//var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	//flag.Parse()
	prometheus.MustRegister(exporter.NewLinuxExporter("port111"))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
