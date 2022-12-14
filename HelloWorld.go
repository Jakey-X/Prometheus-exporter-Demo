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

func recordMetrics() {
	go func() {
		for {
			cc, _ := cpu.Percent(time.Second, false)
			opsProcessed.Set(cc[0])
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewGauge(prometheus.GaugeOpts{
		Name:        "just_a_test",
		Help:        "Nothing is certain",
		ConstLabels: prometheus.Labels{"myid": "jakey"},
	})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
