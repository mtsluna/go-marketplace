package dto

type ImageDto struct {

	Data ImageDataDto `json:"data"`

}

type ImageDataDto struct {

	Link string `json:"link"`

}
