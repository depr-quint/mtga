package inventory

type PlayerInventory struct {
	PlayerId             string      `json:"playerId"`
	CommonWildcards      int         `json:"wcCommon"`
	UncommonWildcards    int         `json:"wcUncommon"`
	RareWildcards        int         `json:"wcRare"`
	MythicWildcards      int         `json:"wcMythic"`
	Gold                 int         `json:"gold"`
	Gems                 int         `json:"gems"`
	DraftTokens          int         `json:"draftTokens"`
	SealedTokens         int         `json:"sealedTokens"`
	WildcardTackPosition int         `json:"wcTrackPosition"`
	VaultProgress        float64     `json:"vaultProgress"`
	Boosters             interface{} `json:"boosters"`
	VanityItems          Items       `json:"vanityItems"`
	Vouchers             interface{} `json:"vouchers"`
	VanitySelections     Selections  `json:"vanitySelections"`
}

type Items struct {
	Pets      []Pet      `json:"pets"`
	Avatars   []Avatar   `json:"avatars"`
	CardBacks []CardBack `json:"cardBacks"`
}

type Pet struct {
	Name string `json:"name"`
	Mods []Mod  `json:"mods"`
}

type Avatar struct {
	Name string      `json:"name"`
	Mods interface{} `json:"mods"`
}

type CardBack struct {
	Name string      `json:"name"`
	Mods interface{} `json:"mods"`
}

type Mod struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Selections struct {
	AvatarSelection   interface{} `json:"avatarSelection"`
	CardBackSelection interface{} `json:"cardBackSelection"`
	PetSelection      interface{} `json:"petSelection"`
	PetModSelections  interface{} `json:"petModSelections"`
}
