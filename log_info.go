package main

import (
	"encoding/json"
	"log"
	"time"
)

type logInfo struct {
	MessageName   logInfoMessage         `json:"messageName"`
	HumanContext  string                 `json:"humanContext"`
	Payload       map[string]interface{} `json:"payloadObject"`
	TransactionId string                 `json:"transactionId"`
}

func (info logInfo) parse() {
	payload, err := json.Marshal(info.Payload)
	if err != nil {
		log.Fatal(err)
	}

	switch info.MessageName {
	case sceneChangeMsg:
		var c sceneChange
		if err := json.Unmarshal(payload, &c); err != nil {
			log.Fatal(err)
		}

	case performanceReportMsg:
		var r performanceReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}

	case bootSequenceReportMsg:
		var r bootSequenceReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}
	}
}

type logInfoMessage string

const (
	sceneChangeMsg        logInfoMessage = "Client.SceneChange"
	performanceReportMsg  logInfoMessage = "Client.PerformanceReport"
	bootSequenceReportMsg logInfoMessage = "Client.BootSequenceReport"
)

// Client changed scenes.
type sceneChange struct {
	From      string    `json:"fromSceneName"`
	To        string    `json:"toSceneName"`
	Timestamp time.Time `json:"timestamp"`
	Duration  string    `json:"duration"`
	Initiator string    `json:"initiator"`
	Context   string    `json:"context"`
	PlayerId  string    `json:"playerId"`
}

// Session Performance Analysis.
type performanceReport struct {
	FrameRateAverage            int     `json:"FrameRateAverage"`
	FrameRateMinimum            int     `json:"FrameRateMinimum"`
	FrameRateDeviation          int     `json:"FrameRateDeviation"`
	FirstFrameRateFrame         int     `json:"FirstFrameRateFrame"`
	LastFrameRateFrame          int     `json:"LastFrameRateFrame"`
	FrameRateOddities           []frame `json:"FrameRateOddities"`
	TotalMemoryAverage          int     `json:"TotalMemoryAverage"`
	TotalMemoryMaximum          int     `json:"TotalMemoryMaximum"`
	TotalMemoryDeviation        int     `json:"TotalMemoryDeviation"`
	FirstTotalMemorySampleFrame int     `json:"FirstTotalMemorySampleFrame"`
	LastTotalMemorySampleFrame  int     `json:"LastTotalMemorySampleFrame"`
	TotalMemoryOddities         interface{} // no example found in log, just an []
	Bookmarks                   interface{} // no example found in log, just an []
	QualitySettings             []qualitySetting `json:"QualitySettings"`
	CpuBenchmark                benchmark        `json:"CpuBenchmark"`
	GpuBenchmark                benchmark        `json:"GpuBenchmark"`
	PlayerId                    string           `json:"playerId"`
}

type frame struct {
	Frame     int `json:"Frame"`
	Value     int `json:"Value"`
	Deviation int `json:"Deviation"`
}

type qualitySetting struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
}

type benchmark struct {
	Median        float64 `json:"Median"`
	Score         float64 `json:"Score"`
	AdjustedScore float64 `json:"AdjustedScore"`
}

// Duration of the application launch process including granular durations of notable events within.
// Durations are in seconds.
type bootSequenceReport struct {
	BuildVersion               string `json:"buildVersion"`
	EndpointHash               string `json:"EndpointHash"`
	HashExistingManifest       string `json:"SimpleAssetDownloader_HashExistingManifest"`
	LoadManifest               string `json:"SimpleAssetDownloader_LoadManifest"`
	DeleteOldOrModifiedFiles   string `json:"SimpleAssetDownloader_DeleteOldOrModifiedFiles"`
	DetermineAssetsToDownload  string `json:"SimpleAssetDownloader_DetermineAssetsToDownload"`
	SimpleAssetDownloader      string `json:"SimpleAssetDownloader"`
	AssetBundles               string `json:"MdnAssetLibraryLoad_AssetBundles"`
	SetupPayloadTypeDictionary string `json:"MdnAssetLibrary_SetupPayloadTypeDictionary"`
	MdnAssetLibraryLoad        string `json:"MdnAssetLibraryLoad"`
	FindObjects                string `json:"MDNGlobals_FindObjects"`
	CreateObjectPool           string `json:"MDNGlobals_CreateObjectPool"`
	CardDatabase               string `json:"MDNGlobals_CardDatabase"`
	MDNGlobals                 string `json:"MDNGlobals"`
	SceneLoad                  string `json:"SceneLoad"`
	ObjectPooling              string `json:"ObjectPooling"`
	SplashScreen               string `json:"SplashScreen"`
	BootSequenceDuration       string `json:"BootSequenceDuration"`
	PlayerId                   string `json:"playerId"`
}
