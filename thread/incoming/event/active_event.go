package event

import (
	"github.com/di-wu/mtga/thread/incoming"
)

type ActiveEvent struct {
	PublicEventName   string `json:"PublicEventName"`
	InternalEventName string `json:"InternalEventName"`
	EventState        string `json:"EventState"`
	EventType         string `json:"EventType"`
	ModuleGlobalData  struct {
		EntryFees []struct {
			CurrencyType string `json:"CurrencyType"`
			Quantity     int    `json:"Quantity"`
			MaxUses      string `json:"MaxUses"`
		} `json:"EntryFees"`
		CollationIds      []int                  `json:"CollationIds"`
		DeckSelect        string                 `json:"DeckSelect"`
		RankUpdateType    string                 `json:"RankUpdateType"`
		MaxWins           int                    `json:"MaxWins"`
		MaxLosses         int                    `json:"MaxLosses"`
		Prizes            []string               `json:"Prizes"`
		ChestDescriptions []incoming.Description `json:"ChestDescriptions"`
	} `json:"ModuleGlobalData"`
	StartTime            string `json:"StartTime"`
	LockedTime           string `json:"LockedTime"`
	ClosedTime           string `json:"ClosedTime"`
	Parameters           string `json:"Parameters"`
	Group                string `json:"Group"`
	PastEntries          string `json:"PastEntries"`
	DisplayPriority      int    `json:"DisplayPriority"`
	IsArenaPlayModeEvent bool   `json:"IsArenaPlayModeEvent"`
	Emblems              string `json:"Emblems"`
	UsePlayerCourse      bool   `json:"UsePlayerCourse"`
	UILayoutOptions      struct {
		ResignBehavior     string `json:"ResignBehavior"`
		WinTrackBehavior   string `json:"WinTrackBehavior"`
		EventBladeBehavior string `json:"EventBladeBehavior"`
		DeckButtonBehavior string `json:"DeckButtonBehavior"`
		TemplateName       string `json:"TemplateName"`
	} `json:"UILayoutOptions"`
	SkipValidation bool `json:"SkipValidation"`
}
