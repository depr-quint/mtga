package client

type ProgressionView struct {
	FromSceneName string `json:"fromSceneName"`
	TrackName     string `json:"trackName"`
	Duration      string `json:"duration"`
	PlayerId      string `json:"playerId"`
}
