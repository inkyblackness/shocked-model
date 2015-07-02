package model

type Texture struct {
	Identifiable

	Properties TextureProperties `json:"properties"`

	Images []Link `json:"images"`
}
