package incoming

type Description struct {
	Image1            string         `json:"image1"`
	Image2            string         `json:"image2"`
	Image3            string         `json:"image3"`
	Prefab            string         `json:"prefab"`
	ReferenceId       string         `json:"referenceId"`
	HeaderLocKey      string         `json:"headerLocKey"`
	DescriptionLocKey string         `json:"descriptionLocKey"`
	Quantity          string         `json:"quantity"`
	LocParams         map[string]int `json:"locParams"`
	AvailableDate     string         `json:"availableDate"`
}