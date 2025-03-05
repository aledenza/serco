package client

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_client_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"clientName", "httpMethod", "methodName", "statusCode"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_client_request_duration_seconds",
			Help: "Duration of HTTP requests.",
		},
		[]string{"clientName", "httpMethod", "methodName"},
	)
	requestErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_client_request_errors_tota",
			Help: "Total number of HTTP requests errors",
		},
		[]string{"clientName", "httpMethod", "methodName"},
	)
)

func clientMetric(clientName, httpMethod, methodName string) func(statusCode int, err error) {
	start := time.Now()
	return func(statusCode int, err error) {
		elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
		requestDuration.WithLabelValues(clientName, httpMethod, methodName).Observe(elapsed)
		requestsTotal.WithLabelValues(clientName, httpMethod, methodName, strconv.Itoa(statusCode)).Inc()
		if err != nil {
			requestErrors.WithLabelValues(clientName, httpMethod, methodName).Inc()
		}
	}
}

func init() {
	prometheus.MustRegister(requestsTotal, requestDuration, requestErrors)
}
