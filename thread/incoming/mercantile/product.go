package mercantile

type Product struct {
	Id                           string `json:"id"`
	Sku                          string `json:"SKU"`
	MaxPurchaseQuantity          int    `json:"MaxPurchaseQuantity"`
	AccountMax                   int    `json:"AccountMax"`
	AccountRemainingFulfillments int    `json:"AccountRemainingFulfillments"`
	Price                        int    `json:"Price"`
	UnitCount                    int    `json:"UnitCount"`
	PurchaseCurrencyType         string `json:"PurchaseCurrencyType"`
	StoreSection                 string `json:"StoreSection"`
	StoreSubSection              string `json:"StoreSubSection"`
	FeaturedIndex                int    `json:"FeaturedIndex"`
	Enabled                      bool   `json:"Enabled"`
}
