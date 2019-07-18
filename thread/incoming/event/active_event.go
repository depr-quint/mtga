package event

import (
	"github.com/di-wu/mtga/thread/incoming"
)

type ActiveEvent struct {
	PublicEventName      string           `json:"PublicEventName"`
	InternalEventName    string           `json:"InternalEventName"`
	EventState           string           `json:"EventState"`
	EventType            string           `json:"EventType"`
	ModuleGlobalData     ModuleGlobalData `json:"ModuleGlobalData"`
	StartTime            string           `json:"StartTime"`
	LockedTime           string           `json:"LockedTime"`
	ClosedTime           string           `json:"ClosedTime"`
	Parameters           interface{}      `json:"Parameters"`
	Group                string           `json:"Group"`
	PastEntries          string           `json:"PastEntries"`
	DisplayPriority      int              `json:"DisplayPriority"`
	IsArenaPlayModeEvent bool             `json:"IsArenaPlayModeEvent"`
	Emblems              interface{}      `json:"Emblems"`
	UsePlayerCourse      bool             `json:"UsePlayerCourse"`
	UILayoutOptions      UILayoutOptions  `json:"UILayoutOptions"`
	SkipValidation       bool             `json:"SkipValidation"`
}

type ModuleGlobalData struct {
	EntryFees         []EntryFee             `json:"EntryFees"`
	CollationIds      []int                  `json:"CollationIds"`
	DeckSelect        string                 `json:"DeckSelect"`
	RankUpdateType    string                 `json:"RankUpdateType"`
	MaxWins           int                    `json:"MaxWins"`
	MaxLosses         int                    `json:"MaxLosses"`
	Prizes            []string               `json:"Prizes"`
	ChestDescriptions []incoming.Description `json:"ChestDescriptions"`
}

type EntryFee struct {
	CurrencyType string      `json:"CurrencyType"`
	Quantity     int         `json:"Quantity"`
	MaxUses      interface{} `json:"MaxUses"`
}

type UILayoutOptions struct {
	ResignBehavior     string `json:"ResignBehavior"`
	WinTrackBehavior   string `json:"WinTrackBehavior"`
	EventBladeBehavior string `json:"EventBladeBehavior"`
	DeckButtonBehavior string `json:"DeckButtonBehavior"`
	TemplateName       string `json:"TemplateName"`
}
