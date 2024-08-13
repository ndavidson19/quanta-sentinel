package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	LogLines = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "log_lines_total",
			Help: "Total number of log lines by container and stream (stdout/stderr)",
		},
		[]string{"container", "stream"},
	)
	
	ProcessingTime = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "processing_time_seconds",
			Help:    "Time taken to process each log line",
			Buckets: prometheus.ExponentialBuckets(0.0001, 2, 10),
		},
		[]string{"container"},
	)
)

func Init() {
	prometheus.MustRegister(LogLines)
	prometheus.MustRegister(ProcessingTime)
}