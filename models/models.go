package models

import "gorm.io/gorm"

type BaseModel struct {
	gorm.Model `json:"-"`
}

type Manifest struct {
	BaseModel
	Name       string      `json:"-"`
	Components []Component `json:"components"`
}

type Component struct {
	BaseModel
	ManifestID    uint   `json:"-"`
	Source        string `json:"source"`
	ComponentType string `json:"component_type"`
	PlaceHolder   string `json:"placeholder"`
	Label         string `json:"label"`
	InputType     string `json:"input_type"`
}
