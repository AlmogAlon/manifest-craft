package services

import (
	"manifest-craft/models"
	"slices"
)

var TypeOptions = map[string][]string{
	"TextField":   {},
	"ComboBox":    {"Read Only", "Read write", "Full Admin"},
	"RadioButton": {},
}

type ComponentService struct{}

func (s *ComponentService) GetOptions(component *models.Component) []string {
	return TypeOptions[component.ComponentType]
}

func (s *ComponentService) isValidValue(component *models.Component, v string) bool {
	options, ok := TypeOptions[component.ComponentType]
	if !ok {
		return false
	}

	if len(options) == 0 {
		return true
	}

	return component.InputType == "String" && slices.Contains(options, v)
}
