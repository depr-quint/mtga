package quest

import (
	"github.com/di-wu/mtga/thread/incoming"
)

type PlayerQuest struct {
	QuestId                string               `json:"questId"`
	Goal                   int                  `json:"goal"`
	LocKey                 string               `json:"locKey"`
	TileResourceId         string               `json:"tileResourceId"`
	TreasureResourceId     string               `json:"treasureResourceId"`
	QuestTrack             string               `json:"questTrack"`
	IsNewQuest             bool                 `json:"isNewQuest"`
	EndingProgress         int                  `json:"endingProgress"`
	StartingProgress       int                  `json:"startingProgress"`
	CanSwap                bool                 `json:"canSwap"`
	InventoryUpdate        string               `json:"inventoryUpdate"`
	ChestDescription       incoming.Description `json:"chestDescription"`
	HoursWaitAfterComplete int                  `json:"hoursWaitAfterComplete"`
}
