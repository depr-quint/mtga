package mercantile

type StoreStatus struct {
	StoreEnabled        bool `json:"StoreEnabled"`
	GemsEnabled         bool `json:"GemsEnabled"`
	PacksEnabled        bool `json:"PacksEnabled"`
	BundlesEnabled      bool `json:"BundlesEnabled"`
	GuildBundlesEnabled bool `json:"GuildBundlesEnabled"`
}
