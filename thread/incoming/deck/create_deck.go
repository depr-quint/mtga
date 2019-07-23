package deck

type CreateDeck struct {
	CardSkins   interface{} `json:"cardSkins"`
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Format      string      `json:"format"`
	DeckTileId  int         `json:"deckTileId"`
	MainDeck    []int       `json:"mainDeck"`
	SideBoard   []int       `json:"sideBoard"`
	CardBack    string      `json:"cardBack"`
	LastUpdated string      `json:"lastUpdated"`
}
