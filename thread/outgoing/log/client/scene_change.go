package client

type SceneChange struct {
	From      string `json:"fromSceneName"`
	To        string `json:"toSceneName"`
	Timestamp string `json:"timestamp"`
	Duration  string `json:"duration"`
	Initiator string `json:"initiator"`
	Context   string `json:"context"`
	PlayerId  string `json:"playerId"`
}
