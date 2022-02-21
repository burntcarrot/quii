package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var PromLoginRequests = promauto.NewCounter(prometheus.CounterOpts{
	Name: "pm_login_requests",
	Help: "The total number of processed events",
})

var PromLoginDurations = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Name:       "pm_login_durations",
		Help:       "Login requests latencies in microseconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

var PromLoginRequestSizes = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:    "pm_login_request_sizes",
		Buckets: []float64{16, 32, 64, 128, 256},
	})
