package inventory

type Format struct {
	Name                 string   `json:"name"`
	Sets                 []string `json:"sets"`
	BannedCards          []string `json:"bannedCards"`
	CardCountRestriction string   `json:"cardCountRestriction"`
}
