package mtga

import (
	"github.com/di-wu/mtga/thread/outgoing"
	"github.com/di-wu/mtga/thread/outgoing/event"
	"github.com/di-wu/mtga/thread/outgoing/inventory"
	"github.com/di-wu/mtga/thread/outgoing/log"
	"github.com/di-wu/mtga/thread/outgoing/log/client"
	"github.com/di-wu/mtga/thread/outgoing/log/duel_scene"
	"github.com/di-wu/mtga/thread/outgoing/quest"
	"os"
	"path/filepath"
	"testing"
)

func TestOutgoing(t *testing.T) {
	filePath := filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt")

	parser := Parser{}
	parser.Outgoing.OnAuthenticate(func(auth outgoing.Authenticate) {})
	parser.Outgoing.OnDeckSubmit(func(deck event.DeckSubmit) {})
	parser.Outgoing.OnGetPlayerCourse(func(event event.Event) {})
	parser.Outgoing.OnJoin(func(event event.Event) {})
	parser.Outgoing.OnJoinQueue(func(queue event.JoinQueue) {})
	parser.Outgoing.OnGetProductCatalog(func(catalog inventory.ProductCatalog) {})
	parser.Outgoing.OnLogError(func(err log.Err) {})
	parser.Outgoing.OnLogInfo(func(info log.Info) {})
	parser.Outgoing.OnBootSequenceReport(func(report client.BootSequenceReport) {})
	parser.Outgoing.OnConnected(func(conn client.Connected) {})
	parser.Outgoing.OnHomeEventNavigation(func(nav client.EventNavigation) {})
	parser.Outgoing.OnInventoryReport(func(report client.InventoryReport) {})
	parser.Outgoing.OnPerformanceReport(func(report client.PerformanceReport) {})
	parser.Outgoing.OnPregameSequenceReport(func(report client.PregameSequenceReport) {})
	parser.Outgoing.OnPurchaseFunnel(func(funnel client.PurchaseFunnel) {})
	parser.Outgoing.OnSceneChange(func(change client.SceneChange) {})
	parser.Outgoing.OnSystemMessageView(func(view client.SystemMessageView) {})
	parser.Outgoing.OnUserDeviceSpecs(func(specs client.UserDeviceSpecs) {})
	parser.Outgoing.OnGameStart(func(start duel_scene.GameStart) {})
	parser.Outgoing.OnGameStop(func(stop duel_scene.GameStop) {})
	parser.Outgoing.OnEndOfMatchReport(func(report duel_scene.EndOfMatchReport) {})
	parser.OnEmotesUsedReport(func(report duel_scene.EmotesUsedReport) {})
	parser.OnGetTrackDetail(func(detail quest.TrackDetail) {})

	// TODO add info log methods

	tail, err := NewTail(filePath)
	if err != nil {
		t.Error(err)
	}

	for l := range tail.Logs() {
		parser.Parse(l)
	}
}
