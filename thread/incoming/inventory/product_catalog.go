package inventory

type ProductCatalog map[string][]Product

type Product struct {
	Ccv         string `json:"ccv"`
	Id          string `json:"id"`
	Currency    string `json:"currency"`
	Amount      int    `json:"amount"`
	Purchasable bool   `json:"purchasable"`
	Free        bool   `json:"free"`
}
