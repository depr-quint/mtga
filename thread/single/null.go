package single

type NullEntity struct {
	Id             int      `json:"id"`
	AffectorId     int      `json:"affectorId"`
	AffectedIds    []int    `json:"affectedIds"`
	Type           []string `json:"type"`
	Details        []Detail `json:"details"`
	AllowRedaction bool     `json:"allowRedaction"`
}

type Detail struct {
	Key        string `json:"key"`
	Type       string `json:"type"`
	ValueInt32 []int  `json:"valueInt32"`
}
