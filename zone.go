package mtga

import (
	"regexp"
	"strconv"
)

type ZoneChange struct {
	Type       ZoneChangeType
	Target     string
	InstanceId int
	GrpId      int
	Instigator int
	Source     string
}

type ZoneChangeType string

const (
	Countered     ZoneChangeType = "Countered"
	Damage        ZoneChangeType = "Damage"
	Destroy       ZoneChangeType = "Destroy"
	Exile         ZoneChangeType = "Exile"
	ZeroToughness ZoneChangeType = "ZeroToughness"
)

func (parser *Parser) OnZoneChange(callback func(change ZoneChange)) {
	parser.onZoneChange = callback
}

func (parser *Parser) parserZoneChange(first string) {
	if parser.onZoneChange != nil {
		str := regexp.MustCompile(`([a-zA-Z]*?) for \[\"([a-zA-Z ,\'-]*?)\" InstanceId:([0-9]*?), GrpId:([0-9]*?)\] \(\".*?\"\) had Instigator ([0-9]*?) \(\"([a-zA-Z ,\'-]*?)\"\)`).FindStringSubmatch(first)
		if str != nil {
			instanceId, _ := strconv.Atoi(str[3])
			grpId, _ := strconv.Atoi(str[4])
			instigator, _ := strconv.Atoi(str[5])
			parser.onZoneChange(ZoneChange{
				Type:       ZoneChangeType(str[1]),
				Target:     str[2],
				InstanceId: instanceId,
				GrpId:      grpId,
				Instigator: instigator,
				Source:     str[6],
			})
		} else {
			null := regexp.MustCompile(`([a-zA-Z]*?) for ([0-9]*?) \(\"\[NULL\]\"\) had Instigator ([0-9]*?) \(\"([a-zA-Z ,\'-]*?)\"\)`).FindStringSubmatch(first)
			instanceId, _ := strconv.Atoi(null[2])
			instigator, _ := strconv.Atoi(null[3])
			parser.onZoneChange(ZoneChange{
				Type:       ZoneChangeType(null[1]),
				Target:     "NULL",
				InstanceId: instanceId,
				Instigator: instigator,
				Source:     null[4],
			})
		}
	}
}
