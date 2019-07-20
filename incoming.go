package mtga

import (
	"encoding/json"
	"fmt"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/incoming"
	"github.com/di-wu/mtga/thread/incoming/deck"
	"github.com/di-wu/mtga/thread/incoming/event"
	"github.com/di-wu/mtga/thread/incoming/front_door"
	"github.com/di-wu/mtga/thread/incoming/inventory"
	"github.com/di-wu/mtga/thread/incoming/mercantile"
	"github.com/di-wu/mtga/thread/incoming/mot_d"
	"github.com/di-wu/mtga/thread/incoming/progression"
	"github.com/di-wu/mtga/thread/incoming/quest"
)

// Incoming is a structure that holds the parser's incoming callbacks.
type Incoming struct {
	// thread/incoming/deck
	onGetDeckLists   func(decks []deck.Deck)
	onGetPreconDecks func(decks []deck.PreconDeck)
	// thread/incoming/event
	onClaimPrize               func(claim event.ClaimPrize)
	onDeckSubmit               func(submit event.DeckSubmit)
	onDrop                     func(course event.Drop)
	onGetActiveEvents          func(events []event.ActiveEvent)
	onGetCombinedRankInfo      func(info event.CombinedRankInfo)
	onGetEventAndSeasonPayouts func(payout event.Payout)
	onGetPlayerCourse          func(course event.Course)
	onGetPlayerCourses         func(courses []event.Course)
	onGetSeasonAndRankDetail   func(detail event.SeasonRankAndDetail)
	onJoin                     func(course event.Course)
	onLeaveQueue               func(leave event.LeaveQueue)
	onPayEntry                 func(entry event.PayEntry)
	// thread/incoming/front_door
	onConnectionDetails func(details front_door.ConnectionDetails)
	// thread/incoming/inventory
	onCrackBooster          func(booster inventory.CrackedBooster)
	onGetCatalogStatus      func(status inventory.CatalogStatus)
	onGetFormats            func(formats []inventory.Format)
	onGetPlayerArtSkins     func(skins inventory.PlayerArtSkins)
	onGetPlayerCards        func(cards inventory.PlayerCards)
	onGetPlayerInventory    func(inventory inventory.PlayerInventory)
	onGetPlayerSequenceData func(data inventory.SequenceData)
	onGetProductCatalog     func(catalog inventory.ProductCatalog)
	onGetRewardSchedule     func(schedule inventory.RewardSchedule)
	// thread/incoming/log
	onLogInfo  func(info []byte)
	onLogError func(message string)
	// thread/incoming/mercantile
	onGetAllProducts func(products []mercantile.Product)
	onGetStoreStatus func(status mercantile.StoreStatus)
	// thread/incoming/mot_d
	onGetMotD func(d mot_d.MotD)
	// thread/incoming/progression
	onGetAllTracks      func(tracks []progression.Track)
	onGetPlayerProgress func(progress progression.PlayerProgress)
	// thread/incoming/quest
	onGetTrackDetail  func(detail quest.TrackDetail)
	onGetPlayerQuests func(quests []quest.PlayerQuest)
	// ?
	onAIPractice           func(success bool)
	onJoinEventQueueStatus func(status bool)
	onJoinQueue            func(success bool)
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

	case incoming.CrackBoosterMethod:
		if parser.Incoming.onCrackBooster != nil {
			var b inventory.CrackedBooster
			err := json.Unmarshal(l.Json, &b)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onCrackBooster(b)
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

	case incoming.ClaimPrizeMethod:
		if parser.Incoming.onClaimPrize != nil {
			var c event.ClaimPrize
			err := json.Unmarshal(l.Json, &c)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onClaimPrize(c)
		}
	case incoming.DeckSubmitMethod:
		if parser.Incoming.onDeckSubmit != nil {
			var s event.DeckSubmit
			err := json.Unmarshal(l.Json, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onDeckSubmit(s)
		}
	case incoming.DropMethod:
		if parser.Incoming.onDrop != nil {
			var d event.Drop
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onDrop(d)
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
	case incoming.GetEventAndSeasonPayoutsMethod:
		if parser.onGetEventAndSeasonPayouts != nil {
			var p event.Payout
			err := json.Unmarshal(l.Json, &p)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetEventAndSeasonPayouts(p)
		}
	case incoming.GetPlayerCourseMethod:
		if parser.Incoming.onGetPlayerCourse != nil {
			var c event.Course
			err := json.Unmarshal(l.Json, &c)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onGetPlayerCourse(c)
		}
	case incoming.GetPlayerCoursesMethod:
		if parser.onGetPlayerCourses != nil {
			var c []event.Course
			err := json.Unmarshal(l.Json, &c)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerCourses(c)
		}
	case incoming.LeaveQueueMethod:
		if parser.onLeaveQueue != nil {
			var q event.LeaveQueue
			err := json.Unmarshal(l.Json, &q)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onLeaveQueue(q)
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
	case incoming.JoinMethod:
		if parser.Incoming.onJoin != nil {
			var e event.Course
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onJoin(e)
		}
	case incoming.PayEntryMethod:
		if parser.Incoming.onPayEntry != nil {
			var e event.PayEntry
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onPayEntry(e)
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
	case incoming.GetPlayerSequenceDataMethod:
		if parser.onGetPlayerSequenceData != nil {
			var d inventory.SequenceData
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerSequenceData(d)
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
	case incoming.GetTrackDetailMethod:
		if parser.Incoming.onGetTrackDetail != nil {
			var d quest.TrackDetail
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Incoming.onGetTrackDetail(d)
		}

	case incoming.AIPracticeMethod:
		if parser.Incoming.onAIPractice != nil {
			parser.Incoming.onAIPractice(string(l.Json) == "Success")
		}
	case incoming.JoinEventQueueStatusMethod:
		if parser.onJoinEventQueueStatus != nil {
			var s bool
			err := json.Unmarshal(l.Json, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onJoinEventQueueStatus(s)
		}
	case incoming.JoinQueueMethod:
		if parser.Incoming.onJoinQueue != nil {
			parser.Incoming.onJoinQueue(string(l.Json) == "Success")
		}

	case thread.LogErrorMethod:
		if parser.Incoming.onLogError != nil {
			parser.Incoming.onLogError(string(l.Json))
		}
	case thread.LogInfoMethod:
		if parser.Incoming.onLogInfo != nil {
			parser.Incoming.onLogInfo(l.Json)
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

// OnClaimPrize attaches the given callback, which will be called on claiming the prize of an event.
func (incoming *Incoming) OnClaimPrize(callback func(claim event.ClaimPrize)) {
	incoming.onClaimPrize = callback
}

// OnDeckSubmit attaches the given callback, which will be called on submitting a deck.
func (incoming *Incoming) OnDeckSubmit(callback func(submit event.DeckSubmit)) {
	incoming.onDeckSubmit = callback
}

// OnDrop attaches the given callback, which will be called on dropping an event.
func (incoming *Incoming) OnDrop(callback func(drop event.Drop)) {
	incoming.onDrop = callback
}

// OnGetActiveEvents attaches the given callback, which will be called on getting all the active events.
func (incoming *Incoming) OnGetActiveEvents(callback func(events []event.ActiveEvent)) {
	incoming.onGetActiveEvents = callback
}

// OnGetCombinedRankInfo attaches the given callback, which will be called on getting the combined rank info.
func (incoming *Incoming) OnGetCombinedRankInfo(callback func(info event.CombinedRankInfo)) {
	incoming.onGetCombinedRankInfo = callback
}

// OnGetEventAndSeasonPayouts attaches the given callback, which will be called on getting the event and season payouts.
func (incoming *Incoming) OnGetEventAndSeasonPayouts(callback func(payout event.Payout)) {
	incoming.onGetEventAndSeasonPayouts = callback
}

// OnGetPlayerCourses attaches the given callback, which will be called on getting the course of the player.
func (incoming *Incoming) OnGetPlayerCourse(callback func(course event.Course)) {
	incoming.onGetPlayerCourse = callback
}

// OnGetPlayerCourses attaches the given callback, which will be called on getting the courses of the player.
func (incoming *Incoming) OnGetPlayerCourses(callback func(courses []event.Course)) {
	incoming.onGetPlayerCourses = callback
}

// OnGetSeasonAndRankDetail attaches the given callback, which will be called on getting the season and rank details.
func (incoming *Incoming) OnGetSeasonAndRankDetail(callback func(detail event.SeasonRankAndDetail)) {
	incoming.onGetSeasonAndRankDetail = callback
}

// OnJoin attaches the given callback, which will be called on joining.
func (incoming *Incoming) OnJoin(callback func(course event.Course)) {
	incoming.onJoin = callback
}

// OnLeaveQueue attaches the given callback, which will be called on leaving the queue.
func (incoming *Incoming) OnLeaveQueue(callback func(leave event.LeaveQueue)) {
	incoming.onLeaveQueue = callback
}

// OnPayEntry attaches the given callback, which will be called on after paying the entry.
func (incoming *Incoming) OnPayEntry(callback func(entry event.PayEntry)) {
	incoming.onPayEntry = callback
}

// OnGetDeckLists attaches the given callback, which will be called on getting the deck lists.
func (incoming *Incoming) OnGetDeckLists(callback func(decks []deck.Deck)) {
	incoming.onGetDeckLists = callback
}

// OnGetPreconDecks attaches the given callback, which will be called on getting the precon deck lists.
func (incoming *Incoming) OnGetPreconDecks(callback func(decks []deck.PreconDeck)) {
	incoming.onGetPreconDecks = callback
}

// OnCrackBooster attaches the given callback, which will be called on getting the cracked booster.
func (incoming *Incoming) OnCrackBooster(callback func(booster inventory.CrackedBooster)) {
	incoming.onCrackBooster = callback
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

// GetPlayerSequenceData attaches the given callback, which will be called on getting the sequence data of the player.
func (incoming *Incoming) OnGetPlayerSequenceData(callback func(data inventory.SequenceData)) {
	incoming.onGetPlayerSequenceData = callback
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

// OnGetTrackDetail attaches the given callback, which will be called on getting the track details.
func (incoming *Incoming) OnGetTrackDetail(callback func(detail quest.TrackDetail)) {
	incoming.onGetTrackDetail = callback
}

// OnAIPractice attaches the given callback, which will be called on getting the ai practice success status.
func (incoming *Incoming) OnAIPractice(callback func(success bool)) {
	incoming.onAIPractice = callback
}

// OnJoinEventQueueStatus attaches the given callback, which will be called on getting the join event queue status.
func (incoming *Incoming) OnJoinEventQueueStatus(callback func(status bool)) {
	incoming.onJoinEventQueueStatus = callback
}

// OnJoinQueue attaches the given callback, which will be called on getting the join queue success status.
func (incoming *Incoming) OnJoinQueue(callback func(success bool)) {
	incoming.onJoinQueue = callback
}

// OnLogInfo attaches the given callback, which will be called on an incoming info log.
func (incoming *Incoming) OnLogInfo(callback func(info []byte)) {
	incoming.onLogInfo = callback
}

func (parser *Parser) parseIncomingLogInfo(l []byte) {
	if string(l) != "True" {
		panic.Fatalf("Unparsed incoming info log: %s", string(l))
	}
}
