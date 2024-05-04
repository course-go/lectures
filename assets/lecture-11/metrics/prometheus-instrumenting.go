package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// START OMIT

func main() {
	processedRequestCounter := promauto.NewCounter(
		prometheus.CounterOpts{Name: "api_http_processed_requests_total"})
	randomNumbersGauge := promauto.NewGaugeVec(
		prometheus.GaugeOpts{Name: "api_random_user_numbers"}, []string{"user_id"})
	go func() {
		for {
			processedRequestCounter.Inc()
			time.Sleep(3 * time.Second)
		}
	}()
	go func() {
		for {
			randomNumbersGauge.WithLabelValues(strconv.Itoa(rand.Int())).Set(rand.Float64())
			time.Sleep(2 * time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

// END OMIT
