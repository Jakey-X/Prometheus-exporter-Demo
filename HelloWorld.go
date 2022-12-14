package main

import (
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Functions to refresh metric values at regular intervals
func recordMetrics() {
	// Open concordance(?or I should say goroutines)
	go func() {
		for {
			// Use cc to accept the CPU utilization returned by the function
			cc, _ := cpu.Percent(time.Second, false) //get CPU utilization based on 1 second
			opsProcessed.Set(cc[0])
			// Refresh every two seconds
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	// Define a metric of type gauge as a global variable.
	// Many of the New* functions under the promauto package are useful for building and automatically registering various types of metrics
	opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		// Options passed to the gauge builder
		Name:        "just_a_test",
		Help:        "Nothing is certain",
		ConstLabels: prometheus.Labels{"myid": "jakey"},
	})
)

func main() {
	// Execute the function and open the concurrent process to refresh the value of the metric every 2 seconds
	recordMetrics()

	// Expose the "/metrics" path to port 8080 and use promhttp to handle http requests, thus implementing exporter
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
