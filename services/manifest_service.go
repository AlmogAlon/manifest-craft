package services

import (
	"manifest-craft/models"
)

type ManifestService struct{}

func (*ManifestService) Validate(m *models.Manifest, payload map[string]string) bool {
	componentService := NewComponentService()

	for _, component := range m.Components {
		value, exists := payload[component.Source]

		if !exists || !componentService.IsValidValue(&component, value) {
			return false
		}
	}

	return true
}
