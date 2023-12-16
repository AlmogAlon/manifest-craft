package models

import "gorm.io/gorm"

type Manifest struct {
	gorm.Model
	Name       string      `json:"name"`
	Components []Component `json:"components"`
}

type Component struct {
	ManifestID    uint   `json:"manifest_id"`
	Source        string `json:"source"`
	ComponentType string `json:"component_type"`
	PlaceHolder   string `json:"placeholder"`
	Label         string `json:"label"`
	InputType     string `json:"input_type"`
}
