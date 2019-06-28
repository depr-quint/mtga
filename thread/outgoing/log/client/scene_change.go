package client

import "time"

type SceneChange struct {
	From      string    `json:"fromSceneName"`
	To        string    `json:"toSceneName"`
	Timestamp time.Time `json:"timestamp"`
	Duration  string    `json:"duration"`
	Initiator string    `json:"initiator"`
	Context   string    `json:"context"`
	PlayerId  string    `json:"playerId"`
}
