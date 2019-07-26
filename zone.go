package mtga

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
