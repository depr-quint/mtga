package inventory

type PlayerInventory struct {
	PlayerId             string  `json:"playerId"`
	CommonWildcards      int     `json:"wcCommon"`
	UncommonWildcards    int     `json:"wcUncommon"`
	RareWildcards        int     `json:"wcRare"`
	MythicWildcards      int     `json:"wcMythic"`
	Gold                 int     `json:"gold"`
	Gems                 int     `json:"gems"`
	DraftTokens          int     `json:"draftTokens"`
	SealedTokens         int     `json:"sealedTokens"`
	WildcardTackPosition int     `json:"wcTrackPosition"`
	VaultProgress        float64 `json:"vaultProgress"`
	Boosters             string  `json:"boosters"`
	VanityItems          struct {
		Pets []struct {
			Name string `json:"name"`
			Mods []Mod  `json:"mods"`
		} `json:"pets"`
		Avatars []struct {
			Name string   `json:"name"`
			Mods []string `json:"mods"`
		} `json:"avatars"`
		CardBacks []struct {
			Name string   `json:"name"`
			Mods []string `json:"mods"`
		} `json:"cardBacks"`
	} `json:"vanityItems"`
	Vouchers         string `json:"vouchers"`
	VanitySelections struct {
		AvatarSelection   string   `json:"avatarSelection"`
		CardBackSelection string   `json:"cardBackSelection"`
		PetSelection      string   `json:"petSelection"`
		PetModSelections  []string `json:"petModSelections"`
	} `json:"vanitySelections"`
}

type Mod struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
