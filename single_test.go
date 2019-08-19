package mtga

import (
	"fmt"
	"github.com/di-wu/mtga/thread/single"
	"reflect"
	"testing"
)

func TestSingleSkinsSeen(t *testing.T) {
	l := RawLog{
		Body: []string{`[UnityCrossThreadLogger]Skins seen: 69853=DA 68740=DA 69706=DA 69706=DA .`},
	}
	var callback1, callback2 bool
	parser := Parser{}
	parser.OnSingleTreadLog(func(log string) {
		callback1 = true
	})
	parser.OnSkinsSeen(func(skins single.Skins) {
		callback2 = true
		if len(skins) != 3 {
			t.Error()
			fmt.Println(skins)
		}
	})
	parser.Parse(l)
	if !callback1 || !callback2 {
		t.Error()
	}
}

func TestSingleCardNotExist(t *testing.T) {
	l := RawLog{
		Body: []string{`[UnityCrossThreadLogger]Card #491 ("Zombie") had ParentId #490 but that card did not exist in the GameState.`},
	}
	var callback bool
	parser := Parser{}
	parser.OnCardNotExist(func(card single.NotExist) {
		callback = true
		if card.CardId != 491 || card.CardName != "Zombie" || card.ParentId != 490 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestSingleNullEntity(t *testing.T) {
	l := RawLog{
		Body: []string{`[UnityCrossThreadLogger]NULL entity on { "id": 2450, "affectorId": 4005, "affectedIds": [ 409 ], "type": [ "AnnotationType_ModifiedToughness", "AnnotationType_ModifiedPower", "AnnotationType_Counter" ], "details": [ { "key": "count", "type": "KeyValuePairValueType_int32", "valueInt32": [ 1 ] }, { "key": "counter_type", "type": "KeyValuePairValueType_int32", "valueInt32": [ 1 ] } ], "allowRedaction": true }`},
	}
	var callback bool
	parser := Parser{}
	parser.OnNullEntity(func(null single.NullEntity) {
		callback = true
		if null.Id != 2450 || null.AffectorId != 4005 || len(null.AffectedIds) != 1 ||
			null.AffectedIds[0] != 409 || null.AllowRedaction == false {
			t.Error()
		}
		if len(null.Type) != 3 || null.Type[0] != "AnnotationType_ModifiedToughness" ||
			null.Type[1] != "AnnotationType_ModifiedPower" || null.Type[2] != "AnnotationType_Counter" {
			t.Error()
		}
		if len(null.Details) != 2 || null.Details[0].Key != "count" || null.Details[1].Key != "counter_type" ||
			null.Details[0].Type != null.Details[1].Type || !reflect.DeepEqual(null.Details[0].ValueInt32, null.Details[1].ValueInt32) {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestSingleStateChange(t *testing.T) {
	l := RawLog{
		Body: []string{`[UnityCrossThreadLogger]STATE CHANGED MatchCompleted -> Disconnected`},
	}
	var callback bool
	parser := Parser{}
	parser.OnStateChange(func(from, to string) {
		callback = true
		if from != "MatchCompleted" || to != "Disconnected" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestUnknownSingleLog(t *testing.T) {
	l := RawLog{
		Body: []string{`[UnityCrossThreadLogger]UnknownLog`},
	}
	var callback bool
	parser := Parser{}
	parser.OnUnknownLog(func(message string) {
		callback = true
		if message != "Unparsed single log: UnknownLog" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
