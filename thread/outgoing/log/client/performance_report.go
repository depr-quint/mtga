package client

type PerformanceReport struct {
	FrameRateAverage            int     `json:"FrameRateAverage"`
	FrameRateMinimum            int     `json:"FrameRateMinimum"`
	FrameRateDeviation          int     `json:"FrameRateDeviation"`
	FirstFrameRateFrame         int     `json:"FirstFrameRateFrame"`
	LastFrameRateFrame          int     `json:"LastFrameRateFrame"`
	FrameRateOddities           []Frame `json:"FrameRateOddities"`
	TotalMemoryAverage          int     `json:"TotalMemoryAverage"`
	TotalMemoryMaximum          int     `json:"TotalMemoryMaximum"`
	TotalMemoryDeviation        int     `json:"TotalMemoryDeviation"`
	FirstTotalMemorySampleFrame int     `json:"FirstTotalMemorySampleFrame"`
	LastTotalMemorySampleFrame  int     `json:"LastTotalMemorySampleFrame"`
	TotalMemoryOddities         []interface{} // no example found in log, just an []
	Bookmarks                   []Bookmark       `json:"Bookmarks"`
	QualitySettings             []QualitySetting `json:"QualitySettings"`
	CpuBenchmark                Benchmark        `json:"CpuBenchmark"`
	GpuBenchmark                Benchmark        `json:"GpuBenchmark"`
	PlayerId                    string           `json:"playerId"`
}

type Frame struct {
	Frame     int `json:"Frame"`
	Value     int `json:"Value"`
	Deviation int `json:"Deviation"`
}

type Bookmark struct {
	Frame int    `json:"Frame"`
	Value string `json:"Value"`
}

type QualitySetting struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
}

type Benchmark struct {
	Median        float64 `json:"Median"`
	Score         float64 `json:"Score"`
	AdjustedScore float64 `json:"AdjustedScore"`
}
