package logparser

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)


var (
	ErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "log_error_count",
			Help: "Number of error logs by container",
		},
		[]string{"container", "error_type"},
	)

	LatencyHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_latency_seconds",
			Help:    "Request latency in seconds",
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 15),
		},
		[]string{"container", "endpoint"},
	)
)

func init() {
	prometheus.MustRegister(ErrorCount)
	prometheus.MustRegister(LatencyHistogram)
}

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
	Container string
	Extra     map[string]interface{}
}

func ParseLog(containerID, logLine string) (*LogEntry, error) {
	var entry LogEntry
	entry.Container = containerID
	entry.Extra = make(map[string]interface{}) // Initialize Extra as an empty map

	// Assume log format: 2006-01-02T15:04:05.000Z [LEVEL] Message (JSON extra fields)
	parts := strings.SplitN(logLine, " ", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid log format")
	}

	// Parse timestamp
	timestamp, err := time.Parse(time.RFC3339Nano, parts[0])
	if err != nil {
		return nil, err
	}
	entry.Timestamp = timestamp

	// Parse level
	entry.Level = strings.Trim(parts[1], "[]")

	// Parse message and extra fields
	messageParts := strings.SplitN(parts[2], " {", 2)
	entry.Message = messageParts[0]

	if len(messageParts) > 1 {
		extraJSON := "{" + messageParts[1]
		err = json.Unmarshal([]byte(extraJSON), &entry.Extra)
		if err != nil {
			return nil, err
		}
	}

	// Update metrics based on log content
	updateMetrics(&entry)

	return &entry, nil
}

func updateMetrics(entry *LogEntry) {
	// Count errors
	if entry.Level == "ERROR" {
		ErrorCount.WithLabelValues(entry.Container, entry.Message).Inc()
	}

	// Parse latency from logs (assuming it's logged)
	if latency, ok := entry.Extra["latency"].(float64); ok {
		if endpoint, ok := entry.Extra["endpoint"].(string); ok {
			LatencyHistogram.WithLabelValues(entry.Container, endpoint).Observe(latency)
		}
	}

	// Add more metric updates based on your specific log format and needs
}
