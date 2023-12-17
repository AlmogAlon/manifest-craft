package storage

import "manifest-craft/models"

var data = []models.Manifest{
	{
		Name: "databaseAccess",
		Components: []models.Component{
			{
				Source:        "reason",
				Label:         "reason",
				ComponentType: "TextField",
				PlaceHolder:   "Enter your Reason for access",
				InputType:     "String",
			},
		},
	},
}

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(name string) *models.Manifest {
	for _, value := range data {
		if value.Name == name {
			result := value
			return &result
		}
	}
	return nil
}
