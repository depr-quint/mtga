package inventory

type WildCardBulk struct {
	IsValid          bool     `json:"isValid"`
	ValidationErrors []string `json:"validationErrors"`
}
