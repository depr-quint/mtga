package minus_one

type RankUpdate struct {
	PlayerId         string `json:"playerId"`
	SeasonOrdinal    int    `json:"seasonOrdinal"`
	NewClass         string `json:"newClass"`
	NewLevel         int    `json:"newLevel"`
	OldLevel         int    `json:"oldLevel"`
	NewStep          int    `json:"newStep"`
	OldStep          int    `json:"oldStep"`
	WasLossProtected bool   `json:"wasLossProtected"`
	RankUpdateType   string `json:"rankUpdateType"`
}
