package draft

type Pick struct {
	DraftId    string `json:"draftId"`
	CardId     string `json:"cardId"`
	PackNumber string `json:"packNumber"`
	PickNumber string `json:"pickNumber"`
}
