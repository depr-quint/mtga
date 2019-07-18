package incoming

import "github.com/di-wu/mtga/thread"

const (
	ConnectionDetailsMethod thread.LogMethod = "FrontDoor.ConnectionDetails"

	GetCatalogStatusMethod         thread.LogMethod = "PlayerInventory.GetCatalogStatus"
	AIPractiveMethod               thread.LogMethod = "Event.AIPractice"
	DeckSubmitMethod               thread.LogMethod = "Event.DeckSubmitV3"
	DropMethod                     thread.LogMethod = "Event.Drop"
	GetActiveEventsMethod          thread.LogMethod = "Event.GetActiveEventsV2"
	GetCombinedRankInfoMethod      thread.LogMethod = "Event.GetCombinedRankInfo"
	GetEventAndSeasonPayoutsMethod thread.LogMethod = "Event.GetEventAndSeasonPayouts"
	GetPlayerCourseMethod          thread.LogMethod = "Event.GetPlayerCourseV2"
	GetPlayerCoursesMethod         thread.LogMethod = "Event.GetPlayerCoursesV2"
	GetSeasonAndRankDetailMethod   thread.LogMethod = "Event.GetSeasonAndRankDetail"
	JoinMethod                     thread.LogMethod = "Event.Join"
	LeaveQueueMethod               thread.LogMethod = "Event.LeaveQueue"
	GetDeckListsMethod             thread.LogMethod = "Deck.GetDeckListsV3"
	GetPreconDeckMethod            thread.LogMethod = "Deck.GetPreconDecks"
	GetFormatsMethod               thread.LogMethod = "PlayerInventory.GetFormats"
	GetPlayerArtSkinsMethod        thread.LogMethod = "PlayerInventory.GetPlayerArtSkins"
	GetPlayerCardsMethod           thread.LogMethod = "PlayerInventory.GetPlayerCardsV3"
	GetPlayerInventoryMethod       thread.LogMethod = "PlayerInventory.GetPlayerInventory"
	GetPlayerSequenceDataMethod    thread.LogMethod = "PlayerInventory.GetPlayerSequenceData"
	GetProductCatalogMethod        thread.LogMethod = "PlayerInventory.GetProductCatalog"
	GetRewardScheduleMethod        thread.LogMethod = "PlayerInventory.GetRewardSchedule"
	GetAllTracksMethod             thread.LogMethod = "Progression.GetAllTracks"
	GetPlayerProgressMethod        thread.LogMethod = "Progression.GetPlayerProgress"
	GetPlayerQuestsMethod          thread.LogMethod = "Quest.GetPlayerQuests"
	GetTrackDetailMethod           thread.LogMethod = "Quest.GetTrackDetail"
	GetAllProductsMethod           thread.LogMethod = "Mercantile.GetAllProducts"
	GetStoreStatusMethod           thread.LogMethod = "Mercantile.GetStoreStatus"
	GetMotDMethod                  thread.LogMethod = "MotD.GetMotD"

	JoinEventQueueStatusMethod thread.LogMethod = "Config.JoinEventQueueStatus"
	JoinQueueMethod            thread.LogMethod = "Event.JoinQueue"
)
