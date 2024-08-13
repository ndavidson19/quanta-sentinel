package alerting

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlertManager(t *testing.T) {
	emailConfig := EmailConfig{
		SMTPHost: "smtp.example.com",
		SMTPPort: 587,
		From:     "alerts@example.com",
		Password: "password123",
	}

	am := NewAlertManager(emailConfig)

	// Test adding an alert
	alert := &Alert{
		Name: "High Error Rate",
		Condition: func() bool {
			return true // Always trigger for testing
		},
		Message:  "Error rate is above threshold",
		Cooldown: 5 * time.Minute,
	}

	am.AddAlert(alert)
	assert.Len(t, am.alerts, 1, "Expected one alert to be added")

	// Test alert triggering
	am.Start()
	time.Sleep(1 * time.Second) // Give some time for the goroutine to run

	// Check if LastAlert was updated
	assert.WithinDuration(t, time.Now(), alert.LastAlert, 2*time.Second, "Expected LastAlert to be updated")
}

func TestAlertManager_sendAlert(t *testing.T) {
	emailConfig := EmailConfig{
		SMTPHost: "smtp.example.com",
		SMTPPort: 587,
		From:     "alerts@example.com",
		Password: "password123",
	}

	am := NewAlertManager(emailConfig)

	alert := &Alert{
		Name:    "Test Alert",
		Message: "Test alert message",
	}

	// This test won't actually send an email, but it will log the attempt
	am.sendAlert(alert)

	// In a real scenario, you might want to mock the email sending functionality
	// and assert that it was called with the correct parameters
}