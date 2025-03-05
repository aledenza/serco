package database

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	databaseSuccessTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_query_success_total",
			Help: "Total number of successful database queries.",
		},
		[]string{"databaseName", "methodName"},
	)
	databaseFailedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_query_failed_total",
			Help: "Total number of failed database queries.",
		},
		[]string{"databaseName", "methodName"},
	)
	databaseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "db_query_duration_seconds",
			Help: "Duration of Database queries.",
		},
		[]string{"databaseName", "methodName"},
	)
)

func databaseMetric(databaseName string, methodName string) func(err error) {
	start := time.Now()
	return func(err error) {
		elapsed := float64(time.Since(start).Nanoseconds()) / 1e9
		databaseDuration.WithLabelValues(databaseName, methodName).Observe(elapsed)
		if err == nil {
			databaseSuccessTotal.WithLabelValues(databaseName, methodName).Inc()
		} else {
			databaseFailedTotal.WithLabelValues(databaseName, methodName).Inc()
		}
	}
}

func init() {
	prometheus.MustRegister(databaseDuration, databaseSuccessTotal, databaseFailedTotal)
}
