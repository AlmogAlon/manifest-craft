package storage

import (
	"manifest-craft/models"
)

var manifests = map[string]uint{
	"databaseAccess": 1,
}

var components = []models.Component{
	{
		ManifestID:    1,
		Source:        "reason",
		Label:         "reason",
		ComponentType: "ComboBox",
		PlaceHolder:   "Enter your Reason for access",
		InputType:     "String",
	},
}

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(name string) *models.Manifest {
	manifestId, exists := manifests[name]

	if !exists {
		return nil
	}

	manifest := &models.Manifest{Name: name}

	for _, value := range components {
		if value.ManifestID == manifestId {
			manifest.Components = append(manifest.Components, value)
		}

	}
	return manifest
}

func (s *MemoryStorage) GetComponent(source string) *models.Component {
	for _, value := range components {
		if value.Source == source {
			result := value
			return &result
		}
	}
	return nil
}
