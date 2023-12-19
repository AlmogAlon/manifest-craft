package component

import (
	"manifest-craft/models"
	"manifest-craft/services/component/providers"
	"slices"
)

type Service struct {
	Providers map[string]providers.Provider
}

func NewComponentService() *Service {
	return &Service{
		Providers: providers.InitProviders(),
	}
}

func (s *Service) GetComponentOptions(component *models.Component) []string {
	p, exists := s.Providers[component.Source]

	if !exists {
		p = &providers.DefaultOptionsProvider{}
	}

	value := *p.GetComponentOptions()

	return value[component.ComponentType]
}

func (s *Service) isValidValue(component *models.Component, v string) bool {
	options := s.GetComponentOptions(component)

	componentIsString := component.InputType == "String"

	if len(options) == 0 {
		return componentIsString
	}

	return componentIsString && slices.Contains(options, v)
}

func (s *Service) Validate(c []models.Component, payload map[string]string) bool {
	for _, c := range c {
		value, exists := payload[c.Source]

		if !exists || !s.isValidValue(&c, value) {
			return false
		}
	}

	return true
}
