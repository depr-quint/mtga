package connect_resp

type Response struct {
	Type          string `json:"type"`
	SystemSeatIds []int  `json:"systemSeatIds"`
	MsgId         int    `json:"msgId"`
	ConnectResp   struct {
		Status      string `json:"status"`
		MajorVer    int    `json:"majorVer"`
		RevisionVer int    `json:"revisionVer"`
		BuildVer    int    `json:"buildVer"`
		ProtoVer    string `json:"protoVer"`
		Settings    struct {
			Stops                                  []Stop `json:"stops"`
			AutoPassOption                         string `json:"autoPassOption"`
			GraveyardOrder                         string `json:"graveyardOrder"`
			ManaSelectionType                      string `json:"manaSelectionType"`
			DefaultAutoPassOption                  string `json:"defaultAutoPassOption"`
			SmartStopsSetting                      string `json:"smartStopsSetting"`
			AutoTapStopsSetting                    string `json:"autoTapStopsSetting"`
			AutoOptionalPaymentCancellationSetting string `json:"autoOptionalPaymentCancellationSetting"`
			TransientStops                         []Stop `json:"transientStops"`
		} `json:"settings"`
		DeckMessage struct {
			DeckCards []int `json:"deckCards"`
		} `json:"deckMessage"`
	} `json:"connectResp"`
}

type Stop struct {
	StopType  string `json:"stopType"`
	AppliesTo string `json:"appliesTo"`
	Status    string `json:"status"`
}
