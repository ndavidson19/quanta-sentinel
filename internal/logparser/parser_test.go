package logparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseLog(t *testing.T) {
	tests := []struct {
		name        string
		containerID string
		logLine     string
		want        LogEntry
		wantErr     bool
	}{
		{
			name:        "Valid log entry",
			containerID: "abc123",
			logLine:     "2023-04-01T12:00:00Z [INFO] This is a log message",
			want: LogEntry{
				Timestamp: time.Date(2023, 4, 1, 12, 0, 0, 0, time.UTC),
				Level:     "INFO",
				Message:   "This is a log message",
				Container: "abc123",
				Extra:     map[string]interface{}{},
			},
			wantErr: false,
		},
		{
			name:        "Invalid timestamp",
			containerID: "def456",
			logLine:     "Invalid timestamp [ERROR] Error message",
			wantErr:     true,
		},
		{
			name:        "With extra fields",
			containerID: "ghi789",
			logLine:     "2023-04-01T12:00:00Z [WARN] Warning message {\"key\": \"value\"}",
			want: LogEntry{
				Timestamp: time.Date(2023, 4, 1, 12, 0, 0, 0, time.UTC),
				Level:     "WARN",
				Message:   "Warning message",
				Container: "ghi789",
				Extra:     map[string]interface{}{"key": "value"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLog(tt.containerID, tt.logLine)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Timestamp, got.Timestamp)
				assert.Equal(t, tt.want.Level, got.Level)
				assert.Equal(t, tt.want.Message, got.Message)
				assert.Equal(t, tt.want.Container, got.Container)
				assert.Equal(t, tt.want.Extra, got.Extra)
			}
		})
	}
}