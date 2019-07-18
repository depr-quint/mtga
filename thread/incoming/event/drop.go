package event

import "github.com/di-wu/mtga/thread/incoming/deck"

type Drop struct {
	Id                 string             `json:"Id"`
	InternalEventName  string             `json:"InternalEventName"`
	PlayerId           string             `json:"PlayerId"`
	ModuleInstanceData ModuleInstanceData `json:"ModuleInstanceData"`
	CurrentEventState  string             `json:"CurrentEventState"`
	CurrentModule      string             `json:"CurrentModule"`
	CardPool           string             `json:"CardPool"`
	CourseDeck         deck.PreconDeck    `json:"CourseDeck"`
	PreviousOpponents  interface{}        `json:"PreviousOpponents"`
}
