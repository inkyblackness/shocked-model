package model

type TextureProperties struct {
	Name       [LanguageCount]*string `json:"name"`
	CantBeUsed [LanguageCount]*string `json:"cantBeUsed"`

	//Climbable bool `json:"climbable"`
}
