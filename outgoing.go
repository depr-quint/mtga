package mtga

import (
	"encoding/json"
	"fmt"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/outgoing"
	"github.com/di-wu/mtga/thread/outgoing/event"
	"github.com/di-wu/mtga/thread/outgoing/inventory"
	"github.com/di-wu/mtga/thread/outgoing/log"
	"github.com/di-wu/mtga/thread/outgoing/log/client"
	"github.com/di-wu/mtga/thread/outgoing/log/duel_scene"
	"github.com/di-wu/mtga/thread/outgoing/quest"
)

// Outgoing is a structure that holds the parser's outgoing callbacks.
type Outgoing struct {
	// thread/outgoing
	onAuthenticate func(auth outgoing.Authenticate)
	// thread/outgoing/event
	onAIPractice      func(practice event.AIPractice)
	onClaimPrize      func(event event.Event)
	onDeckSubmit      func(deck event.DeckSubmit)
	onDrop            func(event event.Event)
	onGetPlayerCourse func(event event.Event)
	onJoinQueue       func(queue event.JoinQueue)
	onJoin            func(event event.Event)
	onPayEntry        func(entry event.PayEntry)
	// thread/outgoing/inventory
	onCrackBooster      func(crack inventory.CrackBooster)
	onGetProductCatalog func(catalog inventory.ProductCatalog)
	// thread/outgoing/log
	onLogError func(err log.Err)
	onLogInfo  func(info log.Info)
	// thread/outgoing/log/client
	onBootSequenceReport         func(report client.BootSequenceReport)
	onConnected                  func(conn client.Connected)
	onHomeEventNavigation        func(nav client.EventNavigation)
	onInventoryReport            func(report client.InventoryReport)
	onPerformanceReport          func(report client.PerformanceReport)
	onPregameSequenceReport      func(report client.PregameSequenceReport)
	onProgressionRewardWebViewed func(view client.RewardWeb)
	onProgressionTrackViewed     func(view client.ProgressionView)
	onPurchaseFunnel             func(funnel client.PurchaseFunnel)
	onSceneChange                func(change client.SceneChange)
	onSetAvatarSelection         func(selection client.AvatarSelection)
	onSystemMessageView          func(view client.SystemMessageView)
	onUserDeviceSpecs            func(specs client.UserDeviceSpecs)
	// thread/outgoing/log/duel_scene
	onGameStart        func(start duel_scene.GameStart)
	onGameStop         func(stop duel_scene.GameStop)
	onEndOfMatchReport func(report duel_scene.EndOfMatchReport)
	onEmotesUsedReport func(report duel_scene.EmotesUsedReport)
	// thread/outgoing/quest
	onGetTrackDetail func(detail quest.TrackDetail)
}

func (parser *Parser) parseOutgoingThreadLog(l thread.Log) {
	switch l.Method {
	case outgoing.AuthenticateMethod:
		if (parser.onAuthenticate) != nil {
			var auth outgoing.Authenticate
			err := json.Unmarshal(l.Json, &auth)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onAuthenticate(auth)
		}

	case outgoing.AIPracticeMethod:
		if parser.Outgoing.onAIPractice != nil {
			var p event.AIPractice
			err := json.Unmarshal(l.Json, &p)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onAIPractice(p)
		}
	case outgoing.ClaimPrizeMethod:
		if parser.Outgoing.onClaimPrize != nil {
			var e event.Event
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onClaimPrize(e)
		}
	case outgoing.DeckSubmitMethod:
		if parser.Outgoing.onDeckSubmit != nil {
			var d event.DeckSubmit
			err := json.Unmarshal(l.Json, &d)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onDeckSubmit(d)
		}
	case outgoing.DropMethod:
		if parser.Outgoing.onDrop != nil {
			var e event.Event
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onDrop(e)
		}
	case outgoing.GetPlayerCourseMethod:
		if parser.Outgoing.onGetPlayerCourse != nil {
			var e event.Event
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onGetPlayerCourse(e)
		}
	case outgoing.JoinMethod:
		if parser.Outgoing.onJoin != nil {
			var e event.Event
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onJoin(e)
		}
	case outgoing.JoinQueueMethod:
		if parser.Outgoing.onJoinQueue != nil {
			var queue event.JoinQueue
			err := json.Unmarshal(l.Json, &queue)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onJoinQueue(queue)
		}
	case outgoing.PayEntryMethod:
		if parser.Outgoing.onPayEntry != nil {
			var entry event.PayEntry
			err := json.Unmarshal(l.Json, &entry)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onPayEntry(entry)
		}

	case outgoing.CrackBoosterMethod:
		if parser.Outgoing.onCrackBooster != nil {
			var crack inventory.CrackBooster
			err := json.Unmarshal(l.Json, &crack)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onCrackBooster(crack)
		}
	case outgoing.GetProductCatalogMethod:
		if parser.Outgoing.onGetProductCatalog != nil {
			var catalog inventory.ProductCatalog
			err := json.Unmarshal(l.Json, &catalog)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onGetProductCatalog(catalog)
		}

	case thread.LogErrorMethod:
		if parser.Outgoing.onLogError != nil {
			var e log.Err
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onLogError(e)
		}
	case thread.LogInfoMethod:
		var info log.Info
		err := json.Unmarshal(l.Json, &info)
		if err != nil {
			panic.Fatalln(err)
		}
		if parser.Outgoing.onLogInfo != nil {
			parser.Outgoing.onLogInfo(info)
		}
		parser.parseOutgoingLogInfo(info)

	case outgoing.TrackDetailMethod:
		if parser.Outgoing.onGetTrackDetail != nil {
			var detail quest.TrackDetail
			err := json.Unmarshal(l.Json, &detail)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.Outgoing.onGetTrackDetail(detail)
		}

	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed outgoing log: %s.\n%s\n", l.Method, l.Json))
		}
	}
}

// OnAuthenticate attaches the given callback, which will be called on authenticating.
func (outgoing *Outgoing) OnAuthenticate(callback func(auth outgoing.Authenticate)) {
	outgoing.onAuthenticate = callback
}

// OnAIPractice attaches the given callback, which will be called on practicing with the AI.
func (outgoing *Outgoing) OnAIPractice(callback func(practice event.AIPractice)) {
	outgoing.onAIPractice = callback
}

// OnClaimPrize attaches the given callback, which will be called on claiming the prize of an event.
func (outgoing *Outgoing) OnClaimPrize(callback func(event event.Event)) {
	outgoing.onClaimPrize = callback
}

// OnJoin attaches the given callback, which will be called on submitting a deck.
func (outgoing *Outgoing) OnDeckSubmit(callback func(deck event.DeckSubmit)) {
	outgoing.onDeckSubmit = callback
}

// OnDrop attaches the given callback, which will be called on dropping an event.
func (outgoing *Outgoing) OnDrop(callback func(event event.Event)) {
	outgoing.onDrop = callback
}

// OnGetPlayerCourse attaches the given callback, which will be called on the request of retrieving the player (v2) courses.
func (outgoing *Outgoing) OnGetPlayerCourse(callback func(event event.Event)) {
	outgoing.onGetPlayerCourse = callback
}

// OnJoin attaches the given callback, which will be called on joining an event.
func (outgoing *Outgoing) OnJoin(callback func(event event.Event)) {
	outgoing.onJoin = callback
}

// OnJoinQueue attaches the given callback, which will be called on joining an event queue.
func (outgoing *Outgoing) OnJoinQueue(callback func(queue event.JoinQueue)) {
	outgoing.onJoinQueue = callback
}

// OnPayEntry attaches the given callback, which will be called on paying the entry.
func (outgoing *Outgoing) OnPayEntry(callback func(entry event.PayEntry)) {
	outgoing.onPayEntry = callback
}

// OnCrackBooster attaches the given callback, which will be called on the request of retrieving a cracked booster.
func (outgoing *Outgoing) OnCrackBooster(callback func(crack inventory.CrackBooster)) {
	outgoing.onCrackBooster = callback
}

// OnGetProductCatalog attaches the given callback, which will be called on the request of retrieving the product catalog.
func (outgoing *Outgoing) OnGetProductCatalog(callback func(catalog inventory.ProductCatalog)) {
	outgoing.onGetProductCatalog = callback
}

// OnLogError attaches the given callback, which will be called on an outgoing error log.
func (outgoing *Outgoing) OnLogError(callback func(err log.Err)) {
	outgoing.onLogError = callback
}

// OnLogInfo attaches the given callback, which will be called on an outgoing info log.
func (outgoing *Outgoing) OnLogInfo(callback func(info log.Info)) {
	outgoing.onLogInfo = callback
}

// OnGetTrackDetail attaches the given callback, which will be called on the request of retrieving the track details.
func (outgoing *Outgoing) OnGetTrackDetail(callback func(detail quest.TrackDetail)) {
	outgoing.onGetTrackDetail = callback
}

func (parser *Parser) parseOutgoingLogInfo(l log.Info) {
	payload, err := json.Marshal(l.Payload)
	if err != nil {
		panic.Fatalln(err)
	}

	switch l.MessageName {
	case log.BootSequenceReportMsg:
		if parser.onBootSequenceReport != nil {
			var r client.BootSequenceReport
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onBootSequenceReport(r)
		}
	case log.ConnectedMsg:
		if parser.onConnected != nil {
			var c client.Connected
			err := json.Unmarshal(payload, &c)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onConnected(c)
		}
	case log.EventNavigationMsg:
		if parser.onHomeEventNavigation != nil {
			var n client.EventNavigation
			err := json.Unmarshal(payload, &n)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onHomeEventNavigation(n)
		}
	case log.InventoryReportMsg:
		if parser.onInventoryReport != nil {
			var r client.InventoryReport
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onInventoryReport(r)
		}
	case log.PerformanceReportMsg:
		if parser.onPerformanceReport != nil {
			var r client.PerformanceReport
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onPerformanceReport(r)
		}
	case log.PregameSequenceReportMsg:
		if parser.onPregameSequenceReport != nil {
			var r client.PregameSequenceReport
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onPregameSequenceReport(r)
		}
	case log.ProgressionRewardWebViewedMsg:
		if parser.onProgressionRewardWebViewed != nil {
			var v client.RewardWeb
			err := json.Unmarshal(payload, &v)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onProgressionRewardWebViewed(v)
		}
	case log.ProgressionTrackViewedMsg:
		if parser.onProgressionTrackViewed != nil {
			var v client.ProgressionView
			err := json.Unmarshal(payload, &v)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onProgressionTrackViewed(v)
		}
	case log.PurchaseFunnelMsg:
		if parser.onPurchaseFunnel != nil {
			var f client.PurchaseFunnel
			err := json.Unmarshal(payload, &f)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onPurchaseFunnel(f)
		}
	case log.SceneChangeMsg:
		if parser.onSceneChange != nil {
			var c client.SceneChange
			err := json.Unmarshal(payload, &c)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onSceneChange(c)
		}
	case log.SetAvatarSelectionMsg:
		if parser.onSetAvatarSelection != nil {
			var s client.AvatarSelection
			err := json.Unmarshal(payload, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onSetAvatarSelection(s)
		}
	case log.SystemMessageViewMsg:
		if parser.onSystemMessageView != nil {
			var v client.SystemMessageView
			err := json.Unmarshal(payload, &v)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onSystemMessageView(v)
		}
	case log.UserDeviceSpecsMsg:
		if parser.onUserDeviceSpecs != nil {
			var s client.UserDeviceSpecs
			err := json.Unmarshal(payload, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onUserDeviceSpecs(s)
		}

	case log.GameStartMsg:
		if parser.onGameStart != nil {
			var s duel_scene.GameStart
			err := json.Unmarshal(payload, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGameStart(s)
		}
	case log.GameStopMsg:
		if parser.onGameStop != nil {
			var s duel_scene.GameStop
			err := json.Unmarshal(payload, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGameStop(s)
		}
	case log.EndOfMatchReportMsg:
		if parser.onEndOfMatchReport != nil {
			var r duel_scene.EndOfMatchReport
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onEndOfMatchReport(r)
		}
	case log.EmotesUsedReportMsg:
		if parser.onEmotesUsedReport != nil {
			var r duel_scene.EmotesUsedReport
			err := json.Unmarshal(payload, &r)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onEmotesUsedReport(r)
		}

	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed outgoing info log: %s.\n%s\n", l.MessageName, l.Payload))
		}
	}
}

// OnBootSequenceReport attaches the given callback, which will be called on the report of the boot sequence.
func (outgoing *Outgoing) OnBootSequenceReport(callback func(report client.BootSequenceReport)) {
	outgoing.onBootSequenceReport = callback
}

// OnConnected attaches the given callback, which will be called on connecting.
func (outgoing *Outgoing) OnConnected(callback func(conn client.Connected)) {
	outgoing.onConnected = callback
}

// OnHomeEventNavigation attaches the given callback, which will be called when the user navigated to an event page from the home page.
func (outgoing *Outgoing) OnHomeEventNavigation(callback func(nav client.EventNavigation)) {
	outgoing.onHomeEventNavigation = callback
}

// OnInventoryReport attaches the given callback, which will be called on the report of the summary of the inventory.
func (outgoing *Outgoing) OnInventoryReport(callback func(report client.InventoryReport)) {
	outgoing.onInventoryReport = callback
}

// OnPerformanceReport attaches the given callback, which will be called on the report of the session performance analysis.
func (outgoing *Outgoing) OnPerformanceReport(callback func(report client.PerformanceReport)) {
	outgoing.onPerformanceReport = callback
}

// OnPregameSequenceReport attaches the given callback, which will be called on te report of the duration of the
// matchmaking processes including granular durations of notable events within. Durations are in seconds.
func (outgoing *Outgoing) OnPregameSequenceReport(callback func(report client.PregameSequenceReport)) {
	outgoing.onPregameSequenceReport = callback
}

// OnProgressionRewardWebViewed attaches the given callback, which will be called on viewing the reward web progression.
func (outgoing *Outgoing) OnProgressionRewardWebViewed(callback func(view client.RewardWeb)) {
	outgoing.onProgressionRewardWebViewed = callback
}

// OnProgressionTrackViewed attaches the given callback, which will be called on viewing the track progression.
func (outgoing *Outgoing) OnProgressionTrackViewed(callback func(view client.ProgressionView)) {
	outgoing.onProgressionTrackViewed = callback
}

// OnPurchaseFunnel attaches the given callback, which will be called on updating available store SKUs.
func (outgoing *Outgoing) OnPurchaseFunnel(callback func(funnel client.PurchaseFunnel)) {
	outgoing.onPurchaseFunnel = callback
}

// OnSceneChange attaches the given callback, which will be called on changing scenes.
func (outgoing *Outgoing) OnSceneChange(callback func(change client.SceneChange)) {
	outgoing.onSceneChange = callback
}

// OnSetAvatarSelection attaches the given callback, which will be called on selecting an avatar.
func (outgoing *Outgoing) OnSetAvatarSelection(callback func(selection client.AvatarSelection)) {
	outgoing.onSetAvatarSelection = callback
}

// OnSystemMessageView attaches the given callback, which will be called on system messages.
func (outgoing *Outgoing) OnSystemMessageView(callback func(view client.SystemMessageView)) {
	outgoing.onSystemMessageView = callback
}

// OnUserDeviceSpecs attaches the given callback, which will be called on the report of the user device specs.
func (outgoing *Outgoing) OnUserDeviceSpecs(callback func(specs client.UserDeviceSpecs)) {
	outgoing.onUserDeviceSpecs = callback
}

// OnGameStart attaches the given callback, which will be called on starting the game within a match.
func (outgoing *Outgoing) OnGameStart(callback func(start duel_scene.GameStart)) {
	outgoing.onGameStart = callback
}

// OnGameStop attaches the given callback, which will be called on ending the game within a match.
func (outgoing *Outgoing) OnGameStop(callback func(stop duel_scene.GameStop)) {
	outgoing.onGameStop = callback
}

// OnEndOfMatchReport attaches the given callback, which will be called on the report of an end of a match.
func (outgoing *Outgoing) OnEndOfMatchReport(callback func(report duel_scene.EndOfMatchReport)) {
	outgoing.onEndOfMatchReport = callback
}

// OnEmotesUsedReport attaches the given callback, which will be called on the report of a tally of emotes used by a player during a match.
func (outgoing *Outgoing) OnEmotesUsedReport(callback func(report duel_scene.EmotesUsedReport)) {
	outgoing.onEmotesUsedReport = callback
}
