package client

type SystemMessageView struct {
	CurrentSceneName string `json:"currentSceneName"`
	Title            string `json:"title"`
	Message          string `json:"message"`
	Timestamp        string `json:"timestamp"`
	Duration         string `json:"duration"`
	PlayerId         string `json:"playerId"`
}
