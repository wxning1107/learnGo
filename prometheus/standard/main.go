package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"learnGoSource/prometheus/standard/utils"
	"log"
	"net/http"
	"time"
)

var (
	targetUrl = "http://localhost:1234/metrics"
)

func main() {
	var listenAddress = flag.String("web.listen-address", ":20195", "Address to listen on for web interface and telemetry.")
	var metricPath = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.Parse()

	fileName := time.Now().Format("2006-01-02") + ".log"
	err, file := utils.OpenFile(fileName)
	if err != nil {
		log.Printf("Open file err: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	http.HandleFunc(*metricPath, metricsHandler)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(targetUrl)
	fmt.Println(resp.Header)
	if err != nil {
		log.Printf("Get %s error: %v", targetUrl, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Wrong atatus code: %d", resp.StatusCode)
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read resp error: %v", err)
	}
	_, err = w.Write(all)
	if err != nil {
		log.Printf("Write to response error: %v", err)
	}
}
