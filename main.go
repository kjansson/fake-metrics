package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	gaugeMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "fake_metric",
			Help: "A fake gauge metric for demonstration purposes.",
		},
		[]string{"name", "site"},
	)
	values = map[string]float64{
		"t1_site1": 10.0,
		"t1_site2": 10.0,
		"t2_site1": 10.0,
		"t2_site2": 10.0,
	}
)

func updateGaugeMetric() {
	for {
		for key := range values {
			oldValue := values[key]
			newValue := oldValue + (rand.Float64()*10 - 5) // Change by up to ±5
			if newValue < 0 {
				newValue = 0
			} else if newValue > 20 {
				newValue = 20
			}
			values[key] = newValue

			labels := splitKey(key)
			gaugeMetric.WithLabelValues(labels[0], labels[1]).Set(newValue)
		}
		time.Sleep(30 * time.Second)
	}
}

func splitKey(key string) []string {
	// Splits the key "t1_site1" into ["t1", "site1"]
	return []string{key[:2], key[3:]}
}

func main() {
	prometheus.MustRegister(gaugeMetric)

	// Start updating the gauge metric in a separate goroutine
	go updateGaugeMetric()

	// Expose metrics endpoint on port 9000
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9000", nil)
}
