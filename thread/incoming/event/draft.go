package event

type Draft struct {
	Id                 string `json:"Id"`
	InternalEventName  string `json:"InternalEventName"`
	PlayerId           string `json:"PlayerId"`
	ModuleInstanceData struct {
		HasPaidEntry string `json:"HasPaidEntry"`
		DraftInfo    struct {
			DraftId string `json:"DraftId"`
		} `json:"DraftInfo"`
	} `json:"ModuleInstanceData"`
	CurrentEventState string   `json:"CurrentEventState"`
	CurrentModule     string   `json:"CurrentModule"`
	CardPool          string   `json:"CardPool"`
	CourseDeck        string   `json:"CourseDeck"`
	PreviousOpponents []string `json:"PreviousOpponents"`
}
