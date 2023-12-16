package storage

import "manifest-craft/models"

var data = map[string][]models.Component{
	"databaseAccess": {
		{
			Source:        "reason",
			Label:         "reason",
			ComponentType: "TextField",
			PlaceHolder:   "Enter your Reason for access",
			InputType:     "String",
		},
	},
}

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id string) *[]models.Component {
	result := data[id]
	return &result
}
