package client

type BootSequenceReport struct {
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
