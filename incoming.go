package mtga

import (
	"encoding/json"
	"fmt"
	"github.com/di-wu/mtga/thread/incoming/mot_d"
	"github.com/di-wu/mtga/thread/incoming/progression"
	"github.com/di-wu/mtga/thread/incoming/quest"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/incoming"
	"github.com/di-wu/mtga/thread/incoming/deck"
	"github.com/di-wu/mtga/thread/incoming/event"
	"github.com/di-wu/mtga/thread/incoming/front_door"
	"github.com/di-wu/mtga/thread/incoming/inventory"
	"github.com/di-wu/mtga/thread/incoming/mercantile"
)

// Incoming is a structure that holds the parser's incoming callbacks.
type Incoming struct {
	// thread/incoming/deck
	onGetDeckLists   func(decks []deck.Deck)
	onGetPreconDecks func(decks []deck.PreconDeck)
	// thread/incoming/event
	onGetActiveEvents        func(events []event.ActiveEvent)
	onGetCombinedRankInfo    func(info event.CombinedRankInfo)
	onGetSeasonAndRankDetail func(detail event.SeasonRankAndDetail)
	// thread/incoming/front_door
	onConnectionDetails func(details front_door.ConnectionDetails)
	// thread/incoming/inventory
	onGetCatalogStatus   func(status inventory.CatalogStatus)
	onGetFormats         func(formats []inventory.Format)
	onGetPlayerArtSkins  func(skins inventory.PlayerArtSkins)
	onGetPlayerCards     func(cards inventory.PlayerCards)
	onGetPlayerInventory func(inventory inventory.PlayerInventory)
	onGetProductCatalog  func(catalog inventory.ProductCatalog)
	onGetRewardSchedule  func(schedule inventory.RewardSchedule)
	// thread/incoming/log
	onIncomingLogInfo func(info []byte)
	// thread/incoming/mercantile
	onGetAllProducts func(products []mercantile.Product)
	onGetStoreStatus func(status mercantile.StoreStatus)
	// thread/incoming/mot_d
	onGetMotD func(d mot_d.MotD)
	// thread/incoming/progression
	onGetAllTracks      func(tracks []progression.Track)
	onGetPlayerProgress func(progress progression.PlayerProgress)
	// thread/incoming/quest
	onGetPlayerQuests func(quests []quest.PlayerQuest)
}

func (parser *Parser) parseIncomingThreadLog(l thread.Log) {
	switch l.Method {
	case incoming.GetDeckListsMethod:
		if parser.onGetDeckLists != nil {
			var d []deck.Deck
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetDeckLists(d)
		}
	case incoming.GetPreconDeckMethod:
		if parser.onGetPreconDecks != nil {
			var d []deck.PreconDeck
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPreconDecks(d)
		}

	case incoming.ConnectionDetailsMethod:
		if parser.onConnectionDetails != nil {
			var d front_door.ConnectionDetails
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onConnectionDetails(d)
		}

	case incoming.GetCatalogStatusMethod:
		if parser.onGetCatalogStatus != nil {
			var s inventory.CatalogStatus
			err := json.Unmarshal(l.Json, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetCatalogStatus(s)
		}

	case incoming.GetActiveEventsMethod:
		if parser.onGetActiveEvents != nil {
			var e []event.ActiveEvent
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetActiveEvents(e)
		}
	case incoming.GetCombinedRankInfoMethod:
		if parser.onGetCombinedRankInfo != nil {
			var i event.CombinedRankInfo
			err := json.Unmarshal(l.Json, &i)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetCombinedRankInfo(i)
		}
	case incoming.GetSeasonAndRankDetailMethod:
		if parser.onGetSeasonAndRankDetail != nil {
			var d event.SeasonRankAndDetail
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetSeasonAndRankDetail(d)
		}

	case incoming.GetFormatsMethod:
		if parser.onGetFormats != nil {
			var f []inventory.Format
			err := json.Unmarshal(l.Json, &f)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetFormats(f)
		}
	case incoming.GetPlayerArtSkinsMethod:
		if parser.onGetPlayerArtSkins != nil {
			var s inventory.PlayerArtSkins
			err := json.Unmarshal(l.Json, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerArtSkins(s)
		}
	case incoming.GetPlayerCardsMethod:
		if parser.onGetPlayerCards != nil {
			var c inventory.PlayerCards
			err := json.Unmarshal(l.Json, &c)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerCards(c)
		}
	case incoming.GetPlayerInventoryMethod:
		if parser.onGetPlayerInventory != nil {
			var i inventory.PlayerInventory
			err := json.Unmarshal(l.Json, &i)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerInventory(i)
		}
	case incoming.GetProductCatalogMethod:
		if parser.Incoming.onGetProductCatalog != nil {
			var c inventory.ProductCatalog
			err := json.Unmarshal(l.Json, &c)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onGetProductCatalog(c)
		}
	case incoming.GetRewardScheduleMethod:
		if parser.onGetRewardSchedule != nil {
			var s inventory.RewardSchedule
			err := json.Unmarshal(l.Json, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetRewardSchedule(s)
		}

	case incoming.GetAllTracksMethod:
		if parser.onGetAllTracks != nil {
			var t []progression.Track
			err := json.Unmarshal(l.Json, &t)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetAllTracks(t)
		}

	case incoming.GetMotDMethod:
		if parser.onGetMotD != nil {
			var d mot_d.MotD
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetMotD(d)
		}

	case incoming.GetPlayerProgressMethod:
		if parser.onGetPlayerProgress != nil {
			var p progression.PlayerProgress
			err := json.Unmarshal(l.Json, &p)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerProgress(p)
		}

	case incoming.GetAllProductsMethod:
		if parser.onGetAllProducts != nil {
			var p []mercantile.Product
			err := json.Unmarshal(l.Json, &p)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetAllProducts(p)
		}
	case incoming.GetStoreStatusMethod:
		if parser.onGetStoreStatus != nil {
			var s mercantile.StoreStatus
			err := json.Unmarshal(l.Json, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetStoreStatus(s)
		}

	case incoming.GetPlayerQuestsMethod:
		if parser.onGetPlayerQuests != nil {
			var q []quest.PlayerQuest
			err := json.Unmarshal(l.Json, &q)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerQuests(q)
		}

	case thread.LogInfoMethod:
		if parser.onIncomingLogInfo != nil {
			parser.onIncomingLogInfo(l.Json)
		}
		parser.parseIncomingLogInfo(l.Json)
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed incoming log: %s.\n%s\n", l.Method, l.Json))
		}
	}
}

// OnConnectionDetails attaches the given callback, which will be called on receiving connection details.
func (incoming *Incoming) OnConnectionDetails(callback func(details front_door.ConnectionDetails)) {
	incoming.onConnectionDetails = callback
}

// OnGetActiveEvents attaches the given callback, which will be called on getting all the active events.
func (incoming *Incoming) OnGetActiveEvents(callback func(events []event.ActiveEvent)) {
	incoming.onGetActiveEvents = callback
}

// OnGetCombinedRankInfo attaches the given callback, which will be called on getting the combined rank info.
func (incoming *Incoming) OnGetCombinedRankInfo(callback func(info event.CombinedRankInfo)) {
	incoming.onGetCombinedRankInfo = callback
}

// OnGetSeasonAndRankDetail attaches the given callback, which will be called on getting the season and rank details.
func (incoming *Incoming) OnGetSeasonAndRankDetail(callback func(detail event.SeasonRankAndDetail)) {
	incoming.onGetSeasonAndRankDetail = callback
}

// OnGetDeckLists attaches the given callback, which will be called on getting the deck lists.
func (incoming *Incoming) OnGetDeckLists(callback func(decks []deck.Deck)) {
	incoming.onGetDeckLists = callback
}

// OnGetPreconDecks attaches the given callback, which will be called on getting the precon deck lists.
func (incoming *Incoming) OnGetPreconDecks(callback func(decks []deck.PreconDeck)) {
	incoming.onGetPreconDecks = callback
}

// OnGetCatalogStatus attaches the given callback, which will be called on getting the catalog status.
func (incoming *Incoming) OnGetCatalogStatus(callback func(status inventory.CatalogStatus)) {
	incoming.onGetCatalogStatus = callback
}

// OnGetFormats attaches the given callback, which will be called on getting the formats.
func (incoming *Incoming) OnGetFormats(callback func(formats []inventory.Format)) {
	incoming.onGetFormats = callback
}

// OnGetPlayerArtSkins attaches the given callback, which will be called on getting the card skins of the player.
func (incoming *Incoming) OnGetPlayerArtSkins(callback func(skins inventory.PlayerArtSkins)) {
	incoming.onGetPlayerArtSkins = callback
}

// OnGetPlayerCards attaches the given callback, which will be called on getting the cards of the player.
func (incoming *Incoming) OnGetPlayerCards(callback func(cards inventory.PlayerCards)) {
	incoming.onGetPlayerCards = callback
}

// OnGetPlayerInventory attaches the given callback, which will be called on getting the inventory of the player.
func (incoming *Incoming) OnGetPlayerInventory(callback func(inventory inventory.PlayerInventory)) {
	incoming.onGetPlayerInventory = callback
}

// OnGetProductCatalog attaches the given callback, which will be called on getting the product catalog.
func (incoming *Incoming) OnGetProductCatalog(callback func(catalog inventory.ProductCatalog)) {
	incoming.onGetProductCatalog = callback
}

// OnGetRewardSchedule attaches the given callback, which will be called on getting the reward schedule.
func (incoming *Incoming) OnGetRewardSchedule(callback func(schedule inventory.RewardSchedule)) {
	incoming.onGetRewardSchedule = callback
}

// OnGetMotD attaches the given callback, which will be called on getting the mot d.
func (incoming *Incoming) OnGetMotD(callback func(d mot_d.MotD)) {
	incoming.onGetMotD = callback
}

// OnGetAllTracks attaches the given callback, which will be called on getting all the tracks.
func (incoming *Incoming) OnGetAllTracks(callback func(tracks []progression.Track)) {
	incoming.onGetAllTracks = callback
}

// OnGetPlayerProgress attaches the given callback, which will be called on getting the progress of the player.
func (incoming *Incoming) OnGetPlayerProgress(callback func(progress progression.PlayerProgress)) {
	incoming.onGetPlayerProgress = callback
}

// OnGetAllProducts attaches the given callback, which will be called on getting all the products.
func (incoming *Incoming) OnGetAllProducts(callback func(products []mercantile.Product)) {
	incoming.onGetAllProducts = callback
}

// OnGetStoreStatus attaches the given callback, which will be called on getting the store status.
func (incoming *Incoming) OnGetStoreStatus(callback func(status mercantile.StoreStatus)) {
	incoming.onGetStoreStatus = callback
}

// OnGetPlayerQuests attaches the given callback, which will be called on getting the quests of the player.
func (incoming *Incoming) OnGetPlayerQuests(callback func(quests []quest.PlayerQuest)) {
	incoming.onGetPlayerQuests = callback
}

// OnLogInfo attaches the given callback, which will be called on an incoming info log.
func (incoming *Incoming) OnLogInfo(callback func(info []byte)) {
	incoming.onIncomingLogInfo = callback
}

func (parser *Parser) parseIncomingLogInfo(l []byte) {
	if string(l) != "True" {
		panic.Fatalf("Unparsed incoming info log: %s", string(l))
	}
}
