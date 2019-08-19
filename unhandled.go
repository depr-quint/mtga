package mtga

import (
	"encoding/json"
	"fmt"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/unhandled"
)

type Unhandled struct {
	onDieRollResults func(results unhandled.DieRollResults)
}

func (parser *Parser) parseUnhandledThreadLog(l thread.Log) {
	switch l.Method {
	case unhandled.DieRollResultsMethod:
		if parser.onDieRollResults != nil {
			results := struct {
				Type               string                   `json:"type"`
				SystemSeatIds      []int                    `json:"systemSeatIds"`
				MsgId              int                      `json:"msgId"`
				DieRollResultsResp unhandled.DieRollResults `json:"dieRollResultsResp"`
			}{}
			err := json.Unmarshal(l.Raw, &results)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onDieRollResults(results.DieRollResultsResp)
		}
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed unhandled log: %s.\n%s", l.Method, l.Raw))
		}
	}
}

func (unhandled *Unhandled) OnDieRollResults(callback func(results unhandled.DieRollResults)) {
	unhandled.onDieRollResults = callback
}
