package inventory

type BasicLandSet struct {
	InventoryDelta        Delta           `json:"inventoryDelta"`
	CardOpenEvents        []CardOpenEvent `json:"cardOpenEvents"`
	HasProgressionRewards bool            `json:"hasProgressionReward"`
	XpGained              int             `json:"xpGained"`
}

type Delta struct {
	CardsAdded         []int         `json:"cardsAdded"`
	BoosterDelta       []interface{} `json:"boosterDelta"`
	DecksAdded         []interface{} `json:"decksAdded"`
	VoucherItemsDelta  []interface{} `json:"voucherItemsDelta"`
	WcCommonDelta      int           `json:"wcCommonDelta"`
	WcUncommonDelta    int           `json:"wcUncommonDelta"`
	WcRareDelta        int           `json:"wcRareDelta"`
	WcMythicDelta      int           `json:"wcMythicDelta"`
	GoldDelta          int           `json:"goldDelta"`
	GemsDelta          int           `json:"gemsDelta"`
	EarnedGemsDelta    int           `json:"earnedGemsDelta"`
	DraftTokensDelta   int           `json:"draftTokensDelta"`
	SealedTokensDelta  int           `json:"sealedTokensDelta"`
	WcTrackPosition    int           `json:"wcTrackPosition"`
	VaultProgressDelta int           `json:"vaultProgressDelta"`
	NewNValCommon      int           `json:"newNValCommon"`
	NewNValUncommon    int           `json:"newNValUncommon"`
	NewNValRare        int           `json:"newNValRare"`
	NewNValMythic      int           `json:"newNValMythic"`
	VanityItemsAdded   []interface{} `json:"vanityItemsAdded"`
	VanityItemsRemoved []interface{} `json:"vanityItemsRemoved"`
	ArtSkinsAdded      []interface{} `json:"artSkinsAdded"`
	BasicLandSet       string        `json:"basicLandSet"`
	InvEtag            string        `json:"invEtag"`
	CardEtag           string        `json:"cardEtag"`
	CosmeticEtag       string        `json:"cosmetic_etag"`
}

type CardOpenEvent struct {
	GrpId       int    `json:"grpId"`
	GoldAwarded int    `json:"goldAwarded"`
	GemsAwarded int    `json:"gemsAwarded"`
	Set         string `json:"set"`
}
