package mtga

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/di-wu/mtga/thread/incoming/deck"
	"github.com/di-wu/mtga/thread/incoming/draft"
	"github.com/di-wu/mtga/thread/incoming/event"
	"github.com/di-wu/mtga/thread/incoming/front_door"
	"github.com/di-wu/mtga/thread/incoming/inventory"
	"github.com/di-wu/mtga/thread/incoming/mercantile"
	"github.com/di-wu/mtga/thread/incoming/mot_d"
	"github.com/di-wu/mtga/thread/incoming/progression"
	"github.com/di-wu/mtga/thread/incoming/quest"
)

func TestIncoming(t *testing.T) {
	filePath := filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt")

	parser := Parser{}
	parser.Incoming.OnConnectionDetails(func(details front_door.ConnectionDetails) {})
	parser.Incoming.OnCreateDeck(func(deck deck.CreateDeck) {})
	parser.Incoming.OnUpdateDeck(func(deck deck.Deck) {})
	parser.Incoming.OnGetDeckLists(func(decks []deck.Deck) {})
	parser.Incoming.OnGetPreconDecks(func(decks []deck.PreconDeck) {})
	parser.Incoming.OnDraftStatus(func(status draft.Status) {})
	parser.Incoming.OnMakePick(func(draft draft.Status) {})
	parser.Incoming.OnGetCatalogStatus(func(status inventory.CatalogStatus) {})
	parser.Incoming.OnClaimPrize(func(claim event.ClaimPrize) {})
	parser.Incoming.OnDeckSubmit(func(submit event.DeckSubmit) {})
	parser.Incoming.OnDrop(func(drop event.Drop) {})
	parser.Incoming.OnDraft(func(draft event.Draft) {})
	parser.Incoming.OnGetActiveEvents(func(events []event.ActiveEvent) {})
	parser.Incoming.OnGetCombinedRankInfo(func(info event.CombinedRankInfo) {})
	parser.Incoming.OnGetEventAndSeasonPayouts(func(payout event.Payout) {})
	parser.Incoming.OnGetPlayerCourses(func(courses []event.Course) {})
	parser.Incoming.OnJoin(func(course event.Course) {})
	parser.Incoming.OnLeaveQueue(func(leave event.LeaveQueue) {})
	parser.Incoming.OnGetSeasonAndRankDetail(func(detail event.SeasonRankAndDetail) {})
	parser.Incoming.OnPayEntry(func(entry event.PayEntry) {})
	parser.Incoming.OnCrackBooster(func(booster inventory.CrackedBooster) {})
	parser.Incoming.OnGetFormats(func(formats []inventory.Format) {})
	parser.Incoming.OnGetPlayerArtSkins(func(skins inventory.PlayerArtSkins) {})
	parser.Incoming.OnGetPlayerCards(func(cards inventory.PlayerCards) {})
	parser.Incoming.OnGetPlayerInventory(func(inventory inventory.PlayerInventory) {})
	parser.Incoming.OnGetPlayerSequenceData(func(data inventory.SequenceData) {})
	parser.Incoming.OnGetProductCatalog(func(catalog inventory.ProductCatalog) {})
	parser.Incoming.OnGetRewardSchedule(func(schedule inventory.RewardSchedule) {})
	parser.Incoming.OnRedeemWildCardBulk(func(redeem inventory.WildCardBulk) {})
	parser.Incoming.OnUpdateBasicLandSet(func(update inventory.BasicLandSet) {})
	parser.Incoming.OnGetMotD(func(d mot_d.MotD) {})
	parser.Incoming.OnGetAllTracks(func(tracks []progression.Track) {})
	parser.Incoming.OnGetPlayerProgress(func(progress progression.PlayerProgress) {})
	parser.Incoming.OnGetAllProducts(func(products []mercantile.Product) {})
	parser.Incoming.OnGetStoreStatus(func(status mercantile.StoreStatus) {})
	parser.Incoming.OnGetPlayerQuests(func(quests []quest.PlayerQuest) {})
	parser.Incoming.OnGetTrackDetail(func(detail quest.TrackDetail) {})
	parser.Incoming.OnAIPractice(func(success bool) {})
	parser.Incoming.OnJoinEventQueueStatus(func(status bool) {})
	parser.Incoming.OnJoinQueue(func(success bool) {})
	parser.Incoming.OnLogInfo(func(info []byte) {})

	tail, err := NewTail(filePath)
	if err != nil {
		t.Error(err)
	}

	for l := range tail.Logs() {
		parser.Parse(l)
	}
}
