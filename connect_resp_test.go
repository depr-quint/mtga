package mtga

import (
	"github.com/di-wu/mtga/thread/connect_resp"
	"reflect"
	"testing"
)

func TestConnectResp(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]ConnectResp {`,
			`	"type": "GREMessageType_ConnectResp",`,
			`	"systemSeatIds": [1],`,
			`	"msgId": 2,`,
			`	"connectResp": {`,
			`		"status": "ConnectionStatus_Success",`,
			`		"majorVer": 1,`,
			`		"revisionVer": 1,`,
			`		"buildVer": 9066,`,
			`		"protoVer": "ProtoVersion_MulliganReq",`,
			`		"settings": {`,
			`			"stops": [`,
			`				{`,
			`					"stopType": "StopType_UpkeepStep",`,
			`					"appliesTo": "SettingScope_Team",`,
			`					"status": "SettingStatus_Clear"`,
			`				}`,
			`			],`,
			`			"autoPassOption": "AutoPassOption_ResolveMyStackEffects",`,
			`			"graveyardOrder": "OrderingType_OrderArbitraryAlways",`,
			`			"manaSelectionType": "ManaSelectionType_Auto",`,
			`			"defaultAutoPassOption": "AutoPassOption_ResolveMyStackEffects",`,
			`			"smartStopsSetting": "SmartStopsSetting_Enable",`,
			`			"autoTapStopsSetting": "AutoTapStopsSetting_Enable",`,
			`			"autoOptionalPaymentCancellationSetting": "Setting_Enable",`,
			`			"transientStops": [`,
			`				{`,
			`					"stopType": "StopType_UpkeepStep",`,
			`					"appliesTo": "SettingScope_Team",`,
			`					"status": "SettingStatus_Clear"`,
			`				}`,
			`			]`,
			`		},`,
			`		"deckMessage": {`,
			`			"deckCards": [`,
			`				1,`,
			`				2,`,
			`				3`,
			`			]`,
			`		}`,
			`	}`,
			`}`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.OnConnectResp(func(resp connect_resp.Response) {
		callback = true
		if resp.MsgId != 2 || resp.ConnectResp.BuildVer != 9066 {
			t.Error()
		}
		if !reflect.DeepEqual(resp.ConnectResp.Settings.Stops, resp.ConnectResp.Settings.TransientStops) {
			t.Error()
		}
		if len(resp.ConnectResp.DeckMessage.DeckCards) != 3 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
