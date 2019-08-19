package mtga

import (
	"encoding/json"
	"fmt"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/minus_one"
)

type MinusOne struct {
	onEventMatchCreated    func(match minus_one.MatchCreated)
	onTrackProgressUpdated func(update []minus_one.TrackProgress)
	onInventoryUpdated     func(update minus_one.InventoryUpdate)
	onRankUpdated          func(update minus_one.RankUpdate)
}

func (parser *Parser) parseMinusOneThreadLog(l thread.Log) {
	switch l.Method {
	case minus_one.EventMatchCreatedMethod:
		if parser.onEventMatchCreated != nil {
			var match minus_one.MatchCreated
			err := json.Unmarshal(l.Raw, &match)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onEventMatchCreated(match)
		}
	case minus_one.TrackProgressUpdatedMethod:
		if parser.onTrackProgressUpdated != nil {
			var update []minus_one.TrackProgress
			err := json.Unmarshal(l.Raw, &update)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onTrackProgressUpdated(update)
		}
	case minus_one.InventoryUpdatedMethod:
		if parser.onInventoryUpdated != nil {
			var update minus_one.InventoryUpdate
			err := json.Unmarshal(l.Raw, &update)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onInventoryUpdated(update)
		}
	case minus_one.RankUpdatedMethod:
		if parser.onRankUpdated != nil {
			var update minus_one.RankUpdate
			err := json.Unmarshal(l.Raw, &update)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onRankUpdated(update)
		}
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed minus one log: %s.\n%s", l.Method, l.Raw))
		}
	}
}

func (one *MinusOne) OnEventMatchCreated(callback func(match minus_one.MatchCreated)) {
	one.onEventMatchCreated = callback
}

func (one *MinusOne) OnTrackProgressUpdated(callback func(update []minus_one.TrackProgress)) {
	one.onTrackProgressUpdated = callback
}

func (one *MinusOne) OnInventoryUpdated(callback func(update minus_one.InventoryUpdate)) {
	one.onInventoryUpdated = callback
}

func (one *MinusOne) OnRankUpdated(callback func(update minus_one.RankUpdate)) {
	one.onRankUpdated = callback
}
