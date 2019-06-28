package deck

type Deck struct {
	CardSkins   []CardSkin `json:"cardSkins"`
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Format      string     `json:"format"`
	DeckTileId  int        `json:"deckTileId"`
	MainDeck    []int      `json:"mainDeck"`
	SideBoard   []int      `json:"sideboard"`
	CardBack    string     `json:"cardBack"`
	LastUpdated string     `json:"lastUpdated"`
}

type CardSkin struct {
	GrpId int    `json:"grpId"`
	Ccv   string `json:"ccv"`
}
