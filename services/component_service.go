package services

import (
	"manifest-craft/models"
	"slices"
)

type ComponentTypes map[string][]string

type OptionProvider interface {
	GetComponentOptions() *ComponentTypes
}

type ComponentService struct {
	Provider map[string]OptionProvider
}

func NewComponentService() *ComponentService {
	provider := map[string]OptionProvider{
		"databaseInstances": &DataBaseInstancesProvider{},
		"databaseRoles":     &DefaultOptionsProvider{},
	}
	return &ComponentService{Provider: provider}
}

func (s *ComponentService) GetOptions(component *models.Component) []string {
	p, exists := s.Provider[component.Source]

	if !exists {
		p = &DefaultOptionsProvider{}
	}

	value := *p.GetComponentOptions()

	return value[component.ComponentType]
}

func (s *ComponentService) isValidValue(component *models.Component, v string) bool {
	options := s.GetOptions(component)

	if len(options) == 0 {
		return component.InputType == "String"
	}

	return component.InputType == "String" && slices.Contains(options, v)
}
