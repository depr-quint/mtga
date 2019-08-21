package mtga

import "testing"

func TestChangeZoneNormal(t *testing.T) {
	l := []string{`<<<<<<<<<< ZoneChange of type Destroy for ["Law-Rune Enforcer" InstanceID:294, GrpID:67740] ("Law-Rune Enforcer") had Instigator 291 ("Tyrant's Scorn").`}

	parser := Parser{}
	parser.OnZoneChange(func(change ZoneChange) {
		if change.Type != Destroy || change.Target != "Law-Rune Enforcer" || change.InstanceID != 294 ||
			change.GrpID != 67740 || change.Instigator != 291 || change.Source != "Tyrant's Scorn" {
			t.Error()
		}
	})
	parser.Parse(l)
}

func TestChangeZoneNull(t *testing.T) {
	l := []string{`<<<<<<<<<< ZoneChange of type ZeroToughness for 338 ("[NULL]") had Instigator 334 ("Cry of the Carnarium").`}

	parser := Parser{}
	parser.OnZoneChange(func(change ZoneChange) {
		if change.Type != ZeroToughness || change.Target != "NULL" || change.InstanceID != 338 ||
			change.GrpID != 0 || change.Instigator != 334 || change.Source != "Cry of the Carnarium" {
			t.Error()
		}
	})
	parser.Parse(l)
}
