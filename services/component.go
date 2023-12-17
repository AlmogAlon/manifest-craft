package services

import "manifest-craft/models"

var TypeOptionsMap = map[string][]string{
	"TextField":   {},
	"ComboBox":    {"Read Only", "Read write", "Full Admin"},
	"RadioButton": {},
}

type ComponentService struct{}

func (s *ComponentService) GetOptions(component *models.Component) []string {
	return TypeOptionsMap[component.ComponentType]
}
