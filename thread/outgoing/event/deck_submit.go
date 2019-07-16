package event

type DeckSubmit struct {
	EventName string `json:"eventName"`
	Deck      string `json:"deck"`
}
