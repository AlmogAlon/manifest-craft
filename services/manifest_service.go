package services

import (
	"manifest-craft/models"
)

type ManifestService struct{}

func (*ManifestService) Validate(m *models.Manifest, payload map[string]string) bool {
	componentService := &ComponentService{}

	for _, component := range m.Components {
		value, exists := payload[component.Source]

		if !exists || !componentService.isValidValue(&component, value) {
			return false
		}
	}

	return true
}
