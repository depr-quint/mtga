package quest

type TrackDetail struct {
	ActiveQuestDetails    []string `json:"ActiveQuestDetails"`
	CompletedQuestDetails []struct {
		LocName    string `json:"LocName"`
		ChainName  string `json:"ChainName"`
		ChainIndex int    `json:"ChainIndex"`
		ChainMax   int    `json:"ChainMax"`
	} `json:"CompletedQuestDetails"`
	DontReplenishBeforeTime string `json:"DontReplenishBeforeTime"`
}
