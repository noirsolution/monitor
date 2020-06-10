package monitor

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Monitor struct {
	Webhook string
	Proxies []string
}

// NewMonitor create a new monitor object, with the Discord webhook link for now
func NewMonitor(webhook string, proxies []string) *Monitor {
	m := &Monitor{
		Webhook: webhook,
		Proxies: proxies,
	}

	return m
}

// SendDiscordWebhook is used to send a new embed to the Monitor's webhook
func (m *Monitor) SendDiscordWebhook(webhook Webhook) error {
	webhookJSON, err := json.Marshal(webhook)

	if err != nil {
		log.Fatal(err)
	}

	_, err = http.Post(m.Webhook, "application/json", bytes.NewBuffer(webhookJSON))

	if err != nil {
		return err
	}

	return nil
}
