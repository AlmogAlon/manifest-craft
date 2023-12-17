package services

import (
	"manifest-craft/models"
	"manifest-craft/services/providers"
	"slices"
)

type ComponentService struct {
	Providers map[string]providers.Provider
}

func initProviders() map[string]providers.Provider {
	return map[string]providers.Provider{
		"databaseInstances": &providers.DataBaseInstancesProvider{},
		"databaseRoles":     &providers.DataBaseRulesProvider{},
	}
}

func NewComponentService() *ComponentService {
	return &ComponentService{
		Providers: initProviders(),
	}
}

func (s *ComponentService) GetComponentOptions(component *models.Component) []string {
	p, exists := s.Providers[component.Source]

	if !exists {
		p = &providers.DefaultOptionsProvider{}
	}

	value := *p.GetComponentOptions()

	return value[component.ComponentType]
}

func (s *ComponentService) IsValidValue(component *models.Component, v string) bool {
	options := s.GetComponentOptions(component)

	if len(options) == 0 {
		return component.InputType == "String"
	}

	return component.InputType == "String" && slices.Contains(options, v)
}
