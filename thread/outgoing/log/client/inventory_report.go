package client

type InventoryReport struct {
	Gold             int    `json:"gold"`
	Gems             int    `json:"gems"`
	WildcardCommon   int    `json:"wcCommon"`
	WildcardUncommon int    `json:"wcUncommon"`
	WildcardRare     int    `json:"wcRare"`
	WildcardMythic   int    `json:"wcMythic"`
	DraftTokens      int    `json:"draftTokens"`
	SealedTokens     int    `json:"sealedTokens"`
	PlayerId         string `json:"playerId"`
}
