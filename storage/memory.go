package storage

import (
	"manifest-craft/models"
)

var manifests = map[string]uint{
	"databaseAccess": 0,
	"cloudAccess":    1,
}

var components = []models.Component{
	{
		ManifestID:    0,
		Source:        "reason",
		Label:         "reason",
		ComponentType: "TextField",
		PlaceHolder:   "Enter your Reason for access",
		InputType:     "String",
	},
	{
		ManifestID:    0,
		Source:        "databaseInstances",
		Label:         "Database Instances",
		ComponentType: "ComboBox",
		PlaceHolder:   "choose database",
		InputType:     "String",
	},
	{
		ManifestID:    0,
		Source:        "databaseRoles",
		Label:         "Database Roles",
		ComponentType: "ComboBox",
		PlaceHolder:   "choose a role",
		InputType:     "String",
	},
	{
		ManifestID:    1,
		Source:        "reason",
		Label:         "reason",
		ComponentType: "TextField",
		PlaceHolder:   "Enter your Reason for access",
		InputType:     "String",
	},
	{
		ManifestID:    1,
		Source:        "accounts",
		Label:         "accounts",
		ComponentType: "MultiComboBox",
		PlaceHolder:   "Enter your Accounts",
		InputType:     "String",
	},
	{
		ManifestID:    1,
		Source:        "user",
		Label:         "user",
		ComponentType: "ComboBox",
		InputType:     "String",
	},
	//{
	//	ManifestID:    1,
	//	Source:        "countries",
	//	Label:         "countries",
	//	ComponentType: "ComboBox",
	//	InputType:     "String",
	//},
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
