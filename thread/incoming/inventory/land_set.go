package inventory

type BasicLandSet struct {
	InventoryDelta struct {
		CardsAdded         []int    `json:"cardsAdded"`
		BoosterDelta       []string `json:"boosterDelta"`
		DecksAdded         []string `json:"decksAdded"`
		VoucherItemsDelta  []string `json:"voucherItemsDelta"`
		WcCommonDelta      int      `json:"wcCommonDelta"`
		WcUncommonDelta    int      `json:"wcUncommonDelta"`
		WcRareDelta        int      `json:"wcRareDelta"`
		WcMythicDelta      int      `json:"wcMythicDelta"`
		GoldDelta          int      `json:"goldDelta"`
		GemsDelta          int      `json:"gemsDelta"`
		EarnedGemsDelta    int      `json:"earnedGemsDelta"`
		DraftTokensDelta   int      `json:"draftTokensDelta"`
		SealedTokensDelta  int      `json:"sealedTokensDelta"`
		WcTrackPosition    int      `json:"wcTrackPosition"`
		VaultProgressDelta int      `json:"vaultProgressDelta"`
		NewNValCommon      int      `json:"newNValCommon"`
		NewNValUncommon    int      `json:"newNValUncommon"`
		NewNValRare        int      `json:"newNValRare"`
		NewNValMythic      int      `json:"newNValMythic"`
		VanityItemsAdded   []string `json:"vanityItemsAdded"`
		VanityItemsRemoved []string `json:"vanityItemsRemoved"`
		ArtSkinsAdded      []string `json:"artSkinsAdded"`
		BasicLandSet       string   `json:"basicLandSet"`
		InvEtag            string   `json:"invEtag"`
		CardEtag           string   `json:"cardEtag"`
		CosmeticEtag       string   `json:"cosmetic_etag"`
	} `json:"inventoryDelta"`
	CardOpenEvents []struct {
		GrpId       int    `json:"grpId"`
		GoldAwarded int    `json:"goldAwarded"`
		GemsAwarded int    `json:"gemsAwarded"`
		Set         string `json:"set"`
	} `json:"cardOpenEvents"`
	HasProgressionRewards bool `json:"hasProgressionReward"`
	XpGained              int  `json:"xpGained"`
}
