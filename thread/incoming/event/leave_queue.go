package event

type LeaveQueue struct {
	DidDrop       bool   `json:"didDrop"`
	ActiveMatches string `json:"activeMatches"`
}
