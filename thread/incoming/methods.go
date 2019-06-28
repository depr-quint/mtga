package incoming

import "github.com/di-wu/mtga/thread"

const (
	ConnectionDetailsMethod   thread.LogMethod = "FrontDoor.ConnectionDetails"
	GetCatalogStatusMethod    thread.LogMethod = "PlayerInventory.GetCatalogStatus"
	GetCombinedRankInfoMethod thread.LogMethod = "Event.GetCombinedRankInfo"
	GetDeckListsMethod        thread.LogMethod = "Deck.GetDeckListsV3"
	GetFormatsMethod          thread.LogMethod = "PlayerInventory.GetFormats"
	GetStoreStatusMethod      thread.LogMethod = "Mercantile.GetStoreStatus"
)
