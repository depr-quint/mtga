package progression

type PlayerProgress struct {
	EepTrack            Progress   `json:"eepTrack"`
	IsEepTrackActive    bool       `json:"isEepTrackActive"`
	ActiveBattlePass    Progress   `json:"activeBattlePass"`
	ExpiredBattlePasses []Progress `json:"expiredBattlePasses"`
}

type Progress struct {
	TrackName              string `json:"trackName"`
	CurrentLevel           int    `json:"currentLevel"`
	CurrentExp             int    `json:"currentExp"`
	CurrentTier            int    `json:"currentTier"`
	CurrentOrbCount        int    `json:"currentOrbCount"`
	UnlockedNodeIds        []int  `json:"unlockedNodeIds"`
	AvailableNodesToUnlock []int  `json:"availableNodesToUnlock"`
}
