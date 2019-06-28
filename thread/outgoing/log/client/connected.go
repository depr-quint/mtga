package client

type Connected struct {
	PlayerId      string   `json:"playerId"`
	ScreenName    string   `json:"screenName"`
	Timestamp     int      `json:"timestamp"`
	ClientVersion string   `json:"clientVersion"`
	Settings      Settings `json:"settings"`
}

type Settings struct {
	Gameplay Gameplay `json:"gameplay"`
	Audio    Audio    `json:"audio"`
	Language Language `json:"language"`
}

type Gameplay struct {
	DisableEmotes                bool `json:"disableEmotes"`
	EvergreenKeywordReminders    bool `json:"evergreenKeywordReminders"`
	AutoTap                      bool `json:"autoTap"`
	AutoOrderTriggeredAbilities  bool `json:"autoOrderTriggeredAbilities"`
	AutoChooseReplacementEffects bool `json:"autoChooseReplacementEffects"`
	ShowPhaseLadder              bool `json:"showPhaseLadder"`
	AllPlayModesToggle           bool `json:"allPlayModesToggle"`
}

type Audio struct {
	Master           float64 `json:"master"`
	Music            float64 `json:"music"`
	Effects          float64 `json:"effects"`
	Voice            float64 `json:"voice"`
	Ambience         float64 `json:"ambience"`
	PlayInBackground bool    `json:"playInBackground"`
}

type Language struct {
	Language string `json:"language"`
}
