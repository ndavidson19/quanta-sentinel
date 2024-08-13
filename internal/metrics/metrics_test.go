package metrics

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestMetricsInitialization(t *testing.T) {
	Init()

	// Test LogLines metric
	if testutil.CollectAndCount(LogLines) == 0 {
		t.Error("LogLines metric not registered")
	}

	// Test ProcessingTime metric
	if testutil.CollectAndCount(ProcessingTime) == 0 {
		t.Error("ProcessingTime metric not registered")
	}
}

func TestLogLinesIncrement(t *testing.T) {
	Init()

	LogLines.WithLabelValues("test-container", "stdout").Inc()

	expected := `
		# HELP log_lines_total Total number of log lines by container and stream (stdout/stderr)
		# TYPE log_lines_total counter
		log_lines_total{container="test-container",stream="stdout"} 1
	`

	err := testutil.CollectAndCompare(LogLines, strings.NewReader(expected))
	if err != nil {
		t.Errorf("Unexpected metric value: %v", err)
	}
}

func TestProcessingTimeObservation(t *testing.T) {
	Init()

	ProcessingTime.WithLabelValues("test-container").Observe(0.1)

	expected := `
		# HELP processing_time_seconds Time taken to process each log line
		# TYPE processing_time_seconds histogram
		processing_time_seconds_bucket{container="test-container",le="0.0001"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0002"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0004"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0008"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0016"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0032"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0064"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0128"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0256"} 0
		processing_time_seconds_bucket{container="test-container",le="0.0512"} 0
		processing_time_seconds_bucket{container="test-container",le="0.1024"} 1
		processing_time_seconds_bucket{container="test-container",le="0.2048"} 1
		processing_time_seconds_bucket{container="test-container",le="0.4096"} 1
		processing_time_seconds_bucket{container="test-container",le="0.8192"} 1
		processing_time_seconds_bucket{container="test-container",le="1.6384"} 1
		processing_time_seconds_bucket{container="test-container",le="3.2768"} 1
		processing_time_seconds_bucket{container="test-container",le="6.5536"} 1
		processing_time_seconds_bucket{container="test-container",le="13.1072"} 1
		processing_time_seconds_bucket{container="test-container",le="26.2144"} 1
		processing_time_seconds_bucket{container="test-container",le="52.4288"} 1
		processing_time_seconds_bucket{container="test-container",le="+Inf"} 1
		processing_time_seconds_sum{container="test-container"} 0.1
		processing_time_seconds_count{container="test-container"} 1
	`

	err := testutil.CollectAndCompare(ProcessingTime, strings.NewReader(expected))
	if err != nil {
		t.Errorf("Unexpected metric value: %v", err)
	}
}