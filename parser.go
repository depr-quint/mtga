package mtga

import (
	"encoding/json"
	panic "log"
	"strings"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/outgoing"
	"github.com/di-wu/mtga/thread/outgoing/event"
	"github.com/di-wu/mtga/thread/outgoing/inventory"
	"github.com/di-wu/mtga/thread/outgoing/log"
	"github.com/di-wu/mtga/thread/outgoing/log/client"
	"github.com/di-wu/mtga/thread/outgoing/log/duel_scene"
	"github.com/di-wu/mtga/thread/outgoing/quest"
)

type Parser struct {
	// thread
	onThreadLog func(log thread.Log)
	// thread/outgoing
	onAuthenticate func(auth outgoing.Authenticate)
	// thread/outgoing/event
	onGetPlayerCourse func(event event.Event)
	onJoinQueue       func(queue event.JoinQueue)
	// thread/outgoing/inventory
	onGetProductCatalog func(catalog inventory.ProductCatalog)
	// thread/outgoing/log
	onLogInfo func(info log.Info)
	// thread/outgoing/log/client
	onBootSequenceReport    func(report client.BootSequenceReport)
	onConnected             func(conn client.Connected)
	onHomeEventNavigation   func(nav client.EventNavigation)
	onInventoryReport       func(report client.InventoryReport)
	onPerformanceReport     func(report client.PerformanceReport)
	onPregameSequenceReport func(report client.PregameSequenceReport)
	onPurchaseFunnel        func(funnel client.PurchaseFunnel)
	onSceneChange           func(change client.SceneChange)
	onUserDeviceSpecs       func(specs client.UserDeviceSpecs)
	// thread/outgoing/log/duel_scene
	onGameStart        func(start duel_scene.GameStart)
	onGameStop         func(stop duel_scene.GameStop)
	onEndOfMatchReport func(report duel_scene.EndOfMatchReport)
	onEmotesUsedReport func(report duel_scene.EmotesUsedReport)
	// thread/outgoing/quest
	onGetTrackDetail func(detail quest.TrackDetail)
}

func (parser *Parser) ParseRawLog(l RawLog) {
	if len(l.body) <= 1 {
		return
	}

	switch first, remaining := l.body[0], l.body[1:]; {
	case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
		threadLog := thread.NewLog(strings.TrimPrefix(first, "[UnityCrossThreadLogger]"), remaining)
		if parser.onThreadLog != nil {
			parser.onThreadLog(threadLog)
		}
		parser.parseTreadLog(threadLog)
	default:
		// log.Fatalf("Unparsed log: %s.\n%s\n", first, remaining)
	}
}

func (parser *Parser) OnTreadLog(callback func(log thread.Log)) {
	parser.onThreadLog = callback
}

func (parser *Parser) parseTreadLog(l thread.Log) {
	if len(l.Json) <= 2 {
		return
	}

	switch l.Type {
	case thread.Outgoing:
		parser.parseOutgoingThreadLog(l)
	default:
		// log.Fatalf("Unparsed log: %s.\n", l.Type)
	}
}

func (parser *Parser) parseOutgoingThreadLog(l thread.Log) {
	switch l.Method {
	case thread.AuthenticateMethod:
		if (parser.onAuthenticate) != nil {
			var auth outgoing.Authenticate
			err := json.Unmarshal(l.Json, &auth)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onAuthenticate(auth)
		}

	case thread.PlayerCourseMethod:
		if parser.onGetPlayerCourse != nil {
			var e event.Event
			err := json.Unmarshal(l.Json, &e)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetPlayerCourse(e)
		}
	case thread.JoinQueueMethod:
		if parser.onJoinQueue != nil {
			var queue event.JoinQueue
			err := json.Unmarshal(l.Json, &queue)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onJoinQueue(queue)
		}

	case thread.ProductCatalogMethod:
		if parser.onGetProductCatalog != nil {
			var catalog inventory.ProductCatalog
			err := json.Unmarshal(l.Json, &catalog)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetProductCatalog(catalog)
		}

	case thread.LogInfoMethod:
		var info log.Info
		err := json.Unmarshal(l.Json, &info)
		if err != nil {
			panic.Fatalln(err)
		}
		if (parser.onLogInfo) != nil {
			parser.onLogInfo(info)
		}
		parser.parseLogInfo(info)

	case thread.TrackDetailMethod:
		if parser.onGetTrackDetail != nil {
			var detail quest.TrackDetail
			err := json.Unmarshal(l.Json, &detail)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetTrackDetail(detail)
		}

	default:
		panic.Fatalf("Unparsed log: %s.\n%s\n", l.Method, l.Json)
	}
}

func (parser *Parser) OnAuthenticate(callback func(auth outgoing.Authenticate)) {
	parser.onAuthenticate = callback
}

func (parser *Parser) OnGetPlayerCourse(callback func(event event.Event)) {
	parser.onGetPlayerCourse = callback
}

func (parser *Parser) OnJoinQueue(callback func(queue event.JoinQueue)) {
	parser.onJoinQueue = callback
}

func (parser *Parser) OnGetProductCatalog(callback func(catalog inventory.ProductCatalog)) {
	parser.onGetProductCatalog = callback
}

func (parser *Parser) OnLogInfo(callback func(info log.Info)) {
	parser.onLogInfo = callback
}

func (parser *Parser) OnGetTrackDetail(callback func(detail quest.TrackDetail)) {
	parser.onGetTrackDetail = callback
}

func (parser *Parser) parseLogInfo(l log.Info) {
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
		panic.Fatalf("Unparsed log: %s.\n%s\n", l.MessageName, l.Payload)
	}
}

func (parser *Parser) OnBootSequenceReport(callback func(report client.BootSequenceReport)) {
	parser.onBootSequenceReport = callback
}

func (parser *Parser) OnConnected(callback func(conn client.Connected)) {
	parser.onConnected = callback
}

func (parser *Parser) OnHomeEventNavigation(callback func(nav client.EventNavigation)) {
	parser.onHomeEventNavigation = callback
}

func (parser *Parser) OnInventoryReport(callback func(report client.InventoryReport)) {
	parser.onInventoryReport = callback
}

func (parser *Parser) OnPerformanceReport(callback func(report client.PerformanceReport)) {
	parser.onPerformanceReport = callback
}

func (parser *Parser) OnPregameSequenceReport(callback func(report client.PregameSequenceReport)) {
	parser.onPregameSequenceReport = callback
}

func (parser *Parser) OnPurchaseFunnel(callback func(funnel client.PurchaseFunnel)) {
	parser.onPurchaseFunnel = callback
}

func (parser *Parser) OnSceneChange(callback func(change client.SceneChange)) {
	parser.onSceneChange = callback
}

func (parser *Parser) OnUserDeviceSpecs(callback func(specs client.UserDeviceSpecs)) {
	parser.onUserDeviceSpecs = callback
}

func (parser *Parser) OnGameStart(callback func(start duel_scene.GameStart)) {
	parser.onGameStart = callback
}

func (parser *Parser) OnGameStop(callback func(stop duel_scene.GameStop)) {
	parser.onGameStop = callback
}

func (parser *Parser) OnEndOfMatchReport(callback func(report duel_scene.EndOfMatchReport)) {
	parser.onEndOfMatchReport = callback
}

func (parser *Parser) OnEmotesUsedReport(callback func(report duel_scene.EmotesUsedReport)) {
	parser.onEmotesUsedReport = callback
}
