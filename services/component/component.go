package component

import (
	log "github.com/sirupsen/logrus"
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

func (s *Service) GetComponentOptions(component *models.Component) ([]string, error) {
	p, exists := s.Providers[component.Source]

	if !exists {
		p = &providers.DefaultOptionsProvider{}
	}

	value, err := p.GetComponentOptions()
	if err != nil {
		log.Error("error getting component options")
		return []string{}, err
	}

	deRef := *value

	return deRef[component.ComponentType], nil
}

func (s *Service) isValidValue(component *models.Component, v string) (bool, error) {
	options, err := s.GetComponentOptions(component)
	if err != nil {
		return false, err
	}

	componentIsString := component.InputType == "String"

	if len(options) == 0 {
		return componentIsString, nil
	}

	return componentIsString && slices.Contains(options, v), nil
}

func (s *Service) Validate(c []models.Component, payload map[string]string) (bool, error) {
	for _, c := range c {
		value, exists := payload[c.Source]
		if !exists {
			return false, nil
		}

		_, err := s.isValidValue(&c, value)

		if err != nil {
			return false, err
		}

	}

	return true, nil
}
