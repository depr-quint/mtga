package incoming

import "github.com/di-wu/mtga/thread"

const (
	ConnectionDetailsMethod   thread.LogMethod = "FrontDoor.ConnectionDetails"
	GetCatalogStatusMethod    thread.LogMethod = "PlayerInventory.GetCatalogStatus"
	GetCombinedRankInfoMethod thread.LogMethod = "Event.GetCombinedRankInfo"
	GetDeckListsMethod        thread.LogMethod = "Deck.GetDeckListsV3"
	GetFormatsMethod          thread.LogMethod = "PlayerInventory.GetFormats"
	GetPlayerArtSkinsMethod   thread.LogMethod = "PlayerInventory.GetPlayerArtSkins"
	GetPlayerCardsMethod      thread.LogMethod = "PlayerInventory.GetPlayerCardsV3"
	GetPlayerInventoryMethod  thread.LogMethod = "PlayerInventory.GetPlayerInventory"
	GetProductCatalogMethod   thread.LogMethod = "PlayerInventory.GetProductCatalog"
	GetRewardScheduleMethod   thread.LogMethod = "PlayerInventory.GetRewardSchedule"
	GetPlayerQuestsMethod     thread.LogMethod = "Quest.GetPlayerQuests"
	GetStoreStatusMethod      thread.LogMethod = "Mercantile.GetStoreStatus"
)
