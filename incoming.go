package mtga

import (
	"encoding/json"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/incoming"
	"github.com/di-wu/mtga/thread/incoming/deck"
	"github.com/di-wu/mtga/thread/incoming/event"
	"github.com/di-wu/mtga/thread/incoming/front_door"
	"github.com/di-wu/mtga/thread/incoming/inventory"
	"github.com/di-wu/mtga/thread/incoming/mercantile"
)

// Outgoing is a structure that holds the parser's incoming callbacks.
type Incoming struct {
	// thread/incoming/deck
	onGetDeckLists func(decks []deck.Deck)
	// thread/incoming/event
	onGetCombinedRankInfo func(info event.CombinedRankInfo)
	// thread/incoming/front_door
	onConnectionDetails func(details front_door.ConnectionDetails)
	// thread/incoming/inventory
	onGetCatalogStatus func(status inventory.CatalogStatus)
	onGetFormats       func(formats []inventory.Format)
	// thread/outgoing/log
	onIncomingLogInfo func(info []byte)
	// thread/incoming/mercantile
	onGetStoreStatus func(status mercantile.StoreStatus)
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
	case incoming.GetCombinedRankInfoMethod:
		if parser.onGetCombinedRankInfo != nil {
			var i event.CombinedRankInfo
			err := json.Unmarshal(l.Json, &i)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetCombinedRankInfo(i)
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

	case incoming.GetStoreStatusMethod:
		if parser.onGetStoreStatus != nil {
			var s mercantile.StoreStatus
			err := json.Unmarshal(l.Json, &s)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onGetStoreStatus(s)
		}

	case thread.LogInfoMethod:
		if parser.onIncomingLogInfo != nil {
			parser.onIncomingLogInfo(l.Json)
		}
		parser.parseIncomingLogInfo(l.Json)
	default:
		// panic.Fatalf("Unparsed incoming log: %s.\n%s\n", l.Method, l.Json)
	}
}

// OnConnectionDetails attaches the given callback, which will be called on receiving connection details.
func (incoming *Incoming) OnConnectionDetails(callback func(details front_door.ConnectionDetails)) {
	incoming.onConnectionDetails = callback
}

// OnGetCombinedRankInfo attaches the given callback, which will be called on getting the combined rank info.
func (incoming *Incoming) OnGetCombinedRankInfo(callback func(info event.CombinedRankInfo)) {
	incoming.onGetCombinedRankInfo = callback
}

// OnGetDeckLists attaches the given callback, which will be called on getting the deck lists.
func (incoming *Incoming) OnGetDeckLists(callback func(decks []deck.Deck)) {
	incoming.onGetDeckLists = callback
}

// OnGetCatalogStatus attaches the given callback, which will be called on getting the catalog status.
func (incoming *Incoming) OnGetCatalogStatus(callback func(status inventory.CatalogStatus)) {
	incoming.onGetCatalogStatus = callback
}

// OnGetFormats attaches the given callback, which will be called on getting the formats.
func (incoming *Incoming) OnGetFormats(callback func(formats []inventory.Format)) {
	incoming.onGetFormats = callback
}

// OnGetStoreStatus attaches the given callback, which will be called on getting the store status.
func (incoming *Incoming) OnGetStoreStatus(callback func(status mercantile.StoreStatus)) {
	incoming.onGetStoreStatus = callback
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
