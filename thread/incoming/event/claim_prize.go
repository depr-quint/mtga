package event

import "github.com/di-wu/mtga/thread/incoming/deck"

type ClaimPrize struct {
	Id                 string `json:"Id"`
	InternalEventName  string `json:"InternalEventName"`
	PlayerId           string `json:"PlayerId"`
	ModuleInstanceData struct {
		HasPaidEntry string `json:"HasPaidEntry"`
		DeckSelected bool   `json:"DeckSelected"`
		WinLossGate  struct {
			MaxWins           int      `json:"MaxWins"`
			MaxLosses         int      `json:"MaxLosses"`
			CurrentWins       int      `json:"CurrentWins"`
			CurrentLosses     int      `json:"CurrentLosses"`
			ProcessedMatchIds []string `json:"ProcessedMatchIds"`
		} `json:"WinLossGate"`
		HasClaimedPrize bool `json:"HasClaimedPrize"`
	} `json:"ModuleInstanceData"`
	CurrentEventState string          `json:"CurrentEventState"`
	CurrentModule     string          `json:"CurrentModule"`
	CardPool          string          `json:"CardPool"`
	CourseDeck        deck.PreconDeck `json:"CourseDeck"`
	PreviousOpponents []string          `json:"PreviousOpponents"`
}
