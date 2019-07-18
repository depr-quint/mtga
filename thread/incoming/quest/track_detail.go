package quest

type TrackDetail struct {
	ActiveQuestDetails      interface{} `json:"ActiveQuestDetails"`
	CompletedQuestDetails   []Detail    `json:"CompletedQuestDetails"`
	DontReplenishBeforeTime string      `json:"DontReplenishBeforeTime"`
}

type Detail struct {
	LocName    string `json:"LocName"`
	ChainName  string `json:"ChainName"`
	ChainIndex int    `json:"ChainIndex"`
	ChainMax   int    `json:"ChainMax"`
}
