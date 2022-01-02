package nft

type NFT struct {
	ImageName string `json:"image-name"`
	GrowthRate string `json:"growth-rate"`
	Name string `json:"name"`
	Signature []byte
	ImageUploadUrl string
}

func LoadFromData()  {
}
