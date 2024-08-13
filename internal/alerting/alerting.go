package alerting

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

type AlertManager struct {
	emailConfig EmailConfig
	alerts      map[string]*Alert
}

type EmailConfig struct {
	SMTPHost string
	SMTPPort int
	From     string
	Password string
}

type Alert struct {
	Name      string
	Condition func() bool
	Message   string
	Cooldown  time.Duration
	LastAlert time.Time
}

func NewAlertManager(emailConfig EmailConfig) *AlertManager {
	return &AlertManager{
		emailConfig: emailConfig,
		alerts:      make(map[string]*Alert),
	}
}

func (am *AlertManager) AddAlert(alert *Alert) {
	am.alerts[alert.Name] = alert
}

func (am *AlertManager) Start() {
	go func() {
		for {
			for _, alert := range am.alerts {
				if alert.Condition() && time.Since(alert.LastAlert) > alert.Cooldown {
					am.sendAlert(alert)
					alert.LastAlert = time.Now()
				}
			}
			time.Sleep(30 * time.Second)
		}
	}()
}

func (am *AlertManager) sendAlert(alert *Alert) {
	log.Printf("Alert triggered: %s", alert.Name)

	m := gomail.NewMessage()
	m.SetHeader("From", am.emailConfig.From)
	m.SetHeader("To", "oncall@example.com")
	m.SetHeader("Subject", fmt.Sprintf("Alert: %s", alert.Name))
	m.SetBody("text/plain", alert.Message)

	d := gomail.NewDialer(am.emailConfig.SMTPHost, am.emailConfig.SMTPPort, am.emailConfig.From, am.emailConfig.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send alert email: %v", err)
	}
}
