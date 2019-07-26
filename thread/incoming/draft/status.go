package draft

type Status struct {
	PlayerId     string  `json:"playerId"`
	EventName    string  `json:"eventName"`
	DraftId      string  `json:"draftId"`
	DraftStatus  string  `json:"draftStatus"`
	PackNumber   int     `json:"packNumber"`
	PickNumber   int     `json:"pickNumber"`
	DraftPack    []int   `json:"draftPack"`
	PickedCards  []int   `json:"pickedCards"`
	RequestUnits float64 `json:"requestUnits"`
}
