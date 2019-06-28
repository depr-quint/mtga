package client

import "time"

type SystemMessageView struct {
	CurrentSceneName string    `json:"currentSceneName"`
	Title            string    `json:"title"`
	Message          string    `json:"message"`
	Timestamp        time.Time `json:"timestamp"`
	Duration         string    `json:"duration"`
	PlayerId         string    `json:"playerId"`
}
