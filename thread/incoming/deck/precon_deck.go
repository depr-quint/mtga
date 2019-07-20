package deck

type PreconDeck struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Format        string `json:"format"`
	ResourceId    string `json:"resourceId"`
	DeckTileId    int    `json:"deckTileId"`
	MainDeck      []Card `json:"mainDeck"`
	Sideboard     []Card `json:"sideboard"`
	LastUpdated   string `json:"lastUpdated"`
	LockedForUse  bool   `json:"lockedForUse"`
	LockedForEdit bool   `json:"lockedForEdit"`
	CardBack      string `json:"cardBack"`
	IsValid       bool   `json:"isValid"`
}

type Card struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}
