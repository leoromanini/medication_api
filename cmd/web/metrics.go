package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	apiRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_requests_total",
			Help: "Total number of API requests received",
		},
		[]string{"path", "method", "status"},
	)

	apiRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "api_request_duration_seconds",
			Help:    "Histogram of response times for API requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(apiRequestCount)
	prometheus.MustRegister(apiRequestDuration)
}
