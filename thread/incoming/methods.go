package incoming

import "github.com/di-wu/mtga/thread"

const (
	ConnectionDetailsMethod      thread.LogMethod = "FrontDoor.ConnectionDetails"
	GetCatalogStatusMethod       thread.LogMethod = "PlayerInventory.GetCatalogStatus"
	GetActiveEventsMethod        thread.LogMethod = "Event.GetActiveEventsV2"
	GetCombinedRankInfoMethod    thread.LogMethod = "Event.GetCombinedRankInfo"
	GetSeasonAndRankDetailMethod thread.LogMethod = "Event.GetSeasonAndRankDetail"
	GetDeckListsMethod           thread.LogMethod = "Deck.GetDeckListsV3"
	GetPreconDeckMethod          thread.LogMethod = "Deck.GetPreconDecks"
	GetFormatsMethod             thread.LogMethod = "PlayerInventory.GetFormats"
	GetPlayerArtSkinsMethod      thread.LogMethod = "PlayerInventory.GetPlayerArtSkins"
	GetPlayerCardsMethod         thread.LogMethod = "PlayerInventory.GetPlayerCardsV3"
	GetPlayerInventoryMethod     thread.LogMethod = "PlayerInventory.GetPlayerInventory"
	GetProductCatalogMethod      thread.LogMethod = "PlayerInventory.GetProductCatalog"
	GetRewardScheduleMethod      thread.LogMethod = "PlayerInventory.GetRewardSchedule"
	GetAllTracksMethod           thread.LogMethod = "Progression.GetAllTracks"
	GetPlayerProgressMethod      thread.LogMethod = "Progression.GetPlayerProgress"
	GetPlayerQuestsMethod        thread.LogMethod = "Quest.GetPlayerQuests"
	GetAllProductsMethod         thread.LogMethod = "Mercantile.GetAllProducts"
	GetStoreStatusMethod         thread.LogMethod = "Mercantile.GetStoreStatus"
	GetMotDMethod                thread.LogMethod = "MotD.GetMotD"
)
