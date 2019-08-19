package minus_one

type InventoryUpdate struct {
	Delta           Delta            `json:"delta"`
	AetherizedCards []AetherizedCard `json:"aetherizedCards"`
	Context         string           `json:"context"`
	XpGained        int              `json:"xpGained"`
}

type Delta struct {
	GemsDelta          int      `json:"gemsDelta"`
	BoosterDelta       []string `json:"boosterDelta"`
	CardsAdded         []string `json:"cardsAdded"`
	DecksAdded         []string `json:"decksAdded"`
	VanityItemsAdded   []string `json:"vanityItemsAdded"`
	VanityItemsRemoved []string `json:"vanityItemsRemoved"`
	DraftTokensDelta   int      `json:"draftTokensDelta"`
	GoldDelta          int      `json:"goldDelta"`
	SealedTokensDelta  int      `json:"sealedTokensDelta"`
	VaultProgressDelta float64  `json:"vaultProgressDelta"`
	WcCommonDelta      int      `json:"wcCommonDelta"`
	WcUncommonDelta    int      `json:"wcUncommonDelta"`
	WcRareDelta        int      `json:"wcRareDelta"`
	WcMythicDelta      int      `json:"wcMythicDelta"`
	ArtSkinsAdded      []string `json:"artSkinsAdded"`
	VoucherItemsDelta  []string `json:"voucherItemsDelta"`
}

type AetherizedCard struct {
	GrpId       int    `json:"grpId"`
	GoldAwarded int    `json:"goldAwarded"`
	GemsAwarded int    `json:"gemsAwarded"`
	Set         string `json:"set"`
}
