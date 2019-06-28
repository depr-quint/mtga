package client

type UserDeviceSpecs struct {
	GraphicsDeviceName          string              `json:"graphicsDeviceName"`
	GraphicsDeviceType          string              `json:"graphicsDeviceType"`
	GraphicsDeviceVendor        string              `json:"graphicsDeviceVendor"`
	GraphicsDeviceVersion       string              `json:"graphicsDeviceVersion"`
	GraphicsMemorySize          int                 `json:"graphicsMemorySize"`
	GraphicsMultiThreaded       bool                `json:"graphicsMultiThreaded"`
	GraphicsShaderLevel         int                 `json:"graphicsShaderLevel"`
	DeviceUniqueIdentifier      string              `json:"deviceUniqueIdentifier"`
	DeviceModel                 string              `json:"deviceModel"`
	DeviceType                  string              `json:"deviceType"`
	OperatingSystem             string              `json:"operatingSystem"`
	OperatingSystemFamily       string              `json:"operatingSystemFamily"`
	ProcessorCount              int                 `json:"processorCount"`
	ProcessorFrequency          int                 `json:"processorFrequency"`
	ProcessorType               string              `json:"processorType"`
	SystemMemorySize            int                 `json:"systemMemorySize"`
	MaxTextureSize              int                 `json:"maxTextureSize"`
	IsWindowed                  bool                `json:"isWindowed"`
	GameResolution              Resolution          `json:"gameResolution"`
	MonitorResolution           Resolution          `json:"monitorResolution"`
	MonitorSupportedResolutions []monitorResolution `json:"monitorSupportedResolutions"`
	PlayerId                    string              `json:"playerId"`
}

type Resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type monitorResolution struct {
	Resolution
	ValidForWindow     bool `json:"validForWindow"`
	ValidForFullscreen bool `json:"validForFullscreens"`
}
