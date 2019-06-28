package duel_scene

type EndOfMatchReport struct {
	MatchId                               string   `json:"matchId"`
	MaxCreatures                          int      `json:"maxCreatures"`
	MaxLands                              int      `json:"maxLands"`
	MaxArtifactsAndEnchantments           int      `json:"maxArtifactsAndEnchantments"`
	LongestPassPriorityWaitTimeInSeconds  string   `json:"longestPassPriorityWaitTimeInSeconds"`
	ShortestPassPriorityWaitTimeInSeconds string   `json:"shortestPassPriorityWaitTimeInSeconds"`
	AveragePassPriorityWaitTimeInSeconds  float64  `json:"averagePassPriorityWaitTimeInSeconds"`
	ReceivedPriorityCount                 int      `json:"receivedPriorityCount"`
	PassedPriorityCount                   int      `json:"passedPriorityCount"`
	RespondedToPriorityCount              int      `json:"respondedToPriorityCount"`
	SpellsCastWithAutoPayCount            int      `json:"spellsCastWithAutoPayCount"`
	SpellsCastWithManualManaCount         int      `json:"spellsCastWithManualManaCount"`
	SpellsCastWithMixedPayManaCount       int      `json:"spellsCastWithMixedPayManaCount"`
	AbilityUseByGrpId                     string   `json:"abilityUseByGrpId"`
	AbilityCanceledByGrpId                string   `json:"abilityCanceledByGrpId"`
	AverageActionsByLocalPhaseStep        string   `json:"averageActionsByLocalPhaseStep"`
	AverageActionsByOpponentPhaseStep     string   `json:"averageActionsByOpponentPhaseStep"`
	InteractionCount                      []string `json:"interactionCount"`
	PlayerId                              string   `json:"playerId"`
}
