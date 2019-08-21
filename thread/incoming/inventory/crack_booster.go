package inventory

type CrackedBooster struct {
	CardsOpened []struct {
		GrpId       int    `json:"grpId"`
		GoldAwarded int    `json:"goldAwarded"`
		GemsAwarded int    `json:"gemsAwarded"`
		Set         string `json:"set"`
	} `json:"cardsOpened"`
	TotalVaultProgress     float64 `json:"totalVaultProgress"`
	WildCardTrackMoves     int     `json:"wildCardTrackMoves"`
	WildCardTrackPosition  int     `json:"wildCardTrackPosition"`
	WildCardTrackCommons   int     `json:"wildCardTrackCommons"`
	WildCardTrackUnCommons int     `json:"wildCardTrackUnCommons"`
	WildCardTrackRares     int     `json:"wildCardTrackRares"`
	WildCardTrackMythics   int     `json:"wildCardTrackMythics"`
}
