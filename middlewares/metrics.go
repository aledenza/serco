package middlewares

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	panicTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "panic_total",
			Help: "Total number of panics.",
		},
		[]string{"methodName"},
	)
	apiRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"httpMethod", "methodName", "statusCode"},
	)
	apiRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests.",
		},
		[]string{"httpMethod", "methodName"},
	)

	externalClientRequestErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "external_client_request_total",
			Help: "Total number of HTTP requests from external clients",
		},
		[]string{"clientName"},
	)
)

func panicMetric(methodName string) {
	panicTotal.WithLabelValues(methodName).Inc()
}

func externalClientMetric(clientName string) {
	externalClientRequestErrors.WithLabelValues(clientName).Inc()
}

func apiMetric() func(httpMethod, methodName string, statusCode int) {
	start := time.Now()
	return func(httpMethod, methodName string, statusCode int) {
		elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
		apiRequestDuration.WithLabelValues(httpMethod, methodName).Observe(elapsed)
		apiRequestsTotal.WithLabelValues(httpMethod, methodName, strconv.Itoa(statusCode)).Inc()
	}
}

func init() {
	prometheus.MustRegister(
		panicTotal,
		apiRequestsTotal,
		apiRequestDuration,
		externalClientRequestErrors,
	)
}
