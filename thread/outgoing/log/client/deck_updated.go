package client

type DeckUpdated struct {
	Context   string `json:"Context"`
	DeckId    string `json:"DeckID"`
	Main      []Card `json:"Main"`
	Sideboard []Card `json:"Sideboard"`
	PlayerId  string `json:"playerId"`
}

type Card struct {
	Card    string `json:"Card"`
	Quality string `json:"Quality"`
}
