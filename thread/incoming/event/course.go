package event

import "github.com/di-wu/mtga/thread/incoming/deck"

type Course struct {
	Id                 string `json:"Id"`
	InternalEventName  string `json:"InternalEventName"`
	PlayerId           string `json:"PlayerId"`
	ModuleInstanceData struct {
		DeckSelected bool `json:"DeckSelected"`
	} `json:"ModuleInstanceData"`
	CurrentEventState string    `json:"CurrentEventState"`
	CurrentModule     string    `json:"CurrentModule"`
	CardPool          string    `json:"CardPool"`
	CourseDeck        deck.Deck `json:"CourseDeck"`
	PreviousOpponents []string  `json:"PreviousOpponents"`
}
