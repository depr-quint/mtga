package inventory

type CatalogStatus struct {
	CatalogsEnabled  bool   `json:"CatalogsEnabled"`
	CardSkinsEnabled bool   `json:"CardSkinsEnabled"`
	AvatarsEnabled   bool   `json:"AvatarsEnabled"`
	CardBacksEnabled bool   `json:"CardBacksEnabled"`
	CatalogHash      string `json:"CatalogHash"`
}
