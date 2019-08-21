package mtga

import (
	"regexp"
	"strconv"
)

// ZoneChange occurs on the changing of zones of a certain target by a certain source.
//
// Two types:
//	[Type] for ["[Target]" InstanceID:[ID], GrpID:[ID]] ("[Target]") had Instigator [ID] ("[Source]")
//	[Type] for [ID] ("[NULL]") had Instigator [ID] ("[Source]")
type ZoneChange struct {
	Type       ZoneChangeType
	Target     string
	InstanceID int
	GrpID      int
	Instigator int
	Source     string
}

// ZoneChangeType is the cause type of the zone change.
type ZoneChangeType string

const (
	// Countered indicates that the zone change is caused by a counter event.
	Countered ZoneChangeType = "Countered"
	// Damage indicated that the zone change is caused by a damage event.
	Damage ZoneChangeType = "Damage"
	// Destroy indicated that the zone change is caused by a destroy event.
	Destroy ZoneChangeType = "Destroy"
	// Exile indicated that the zone change is caused by a exile event.
	Exile ZoneChangeType = "Exile"
	// ZeroToughness indicated that the zone change is caused by a zero toughness event.
	ZeroToughness ZoneChangeType = "ZeroToughness"
)

// OnZoneChange attaches the given callback, which will be called on changing zones.
func (parser *Parser) OnZoneChange(callback func(change ZoneChange)) {
	parser.onZoneChange = callback
}

func (parser *Parser) parserZoneChange(first string) {
	if parser.onZoneChange != nil {
		str := regexp.MustCompile(`([a-zA-Z]*?) for \[\"([a-zA-Z ,\'-]*?)\" InstanceID:([0-9]*?), GrpID:([0-9]*?)\] \(\".*?\"\) had Instigator ([0-9]*?) \(\"([a-zA-Z ,\'-]*?)\"\)`).FindStringSubmatch(first)
		if str != nil {
			instanceID, _ := strconv.Atoi(str[3])
			grpID, _ := strconv.Atoi(str[4])
			instigator, _ := strconv.Atoi(str[5])
			parser.onZoneChange(ZoneChange{
				Type:       ZoneChangeType(str[1]),
				Target:     str[2],
				InstanceID: instanceID,
				GrpID:      grpID,
				Instigator: instigator,
				Source:     str[6],
			})
		} else {
			null := regexp.MustCompile(`([a-zA-Z]*?) for ([0-9]*?) \(\"\[NULL\]\"\) had Instigator ([0-9]*?) \(\"([a-zA-Z ,\'-]*?)\"\)`).FindStringSubmatch(first)
			instanceID, _ := strconv.Atoi(null[2])
			instigator, _ := strconv.Atoi(null[3])
			parser.onZoneChange(ZoneChange{
				Type:       ZoneChangeType(null[1]),
				Target:     "NULL",
				InstanceID: instanceID,
				Instigator: instigator,
				Source:     null[4],
			})
		}
	}
}
