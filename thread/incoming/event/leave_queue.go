package event

type LeaveQueue struct {
	DidDrop       bool        `json:"didDrop"`
	ActiveMatches interface{} `json:"activeMatches"`
}
