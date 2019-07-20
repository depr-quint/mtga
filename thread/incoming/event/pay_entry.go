package event

type PayEntry struct {
	Id                 string                  `json:"Id"`
	InternalEventName  string                  `json:"InternalEventName"`
	PlayerId           string                  `json:"PlayerId"`
	ModuleInstanceData EntryModuleInstanceData `json:"ModuleInstanceData"`
	CurrentEventState  string                  `json:"CurrentEventState"`
	CurrentModule      string                  `json:"CurrentModule"`
	CardPool           string                  `json:"CardPool"`
	CourseDecks        string                  `json:"CourseDeck"`
	PreviousOpponents  interface{}             `json:"PreviousOpponents"`
}

type EntryModuleInstanceData struct {
	HasPaidEntry string `json:"HasPaidEntry"`
}
