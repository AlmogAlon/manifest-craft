package component

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"manifest-craft/models"
	"manifest-craft/services/component/providers"
	"reflect"
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

func (s *Service) isValidValue(component *models.Component, v interface{}) (bool, error) {
	x := reflect.ValueOf(v)

	switch x.Kind() {
	case reflect.Bool:
		return component.ComponentType == "RadioButton" && component.InputType == "Boolean", nil

	case reflect.String:
		if component.InputType != "String" {
			log.Error(fmt.Sprintf("got string value for %s component %s", component.InputType, component.Source))
			return false, fmt.Errorf(fmt.Sprintf("wrong value for %s source", component.Source))
		}

		if component.ComponentType == "TextField" {
			return true, nil
		}

		if component.ComponentType != "ComboBox" {
			log.Error(fmt.Sprintf("got string value for %s component %s", component.InputType, component.Source))
			return false, fmt.Errorf(fmt.Sprintf("wrong value for %s source", component.Source))
		}

		options, err := s.GetComponentOptions(component)
		if err != nil {
			return false, err
		}

		if len(options) == 0 {
			return true, nil
		}

		return slices.Contains(options, x.String()), nil

	case reflect.Slice:

		options, err := s.GetComponentOptions(component)
		if err != nil {
			return false, err
		}

		if len(options) == 0 {
			return true, nil
		}

		for i := 0; i < x.Len(); i++ {
			value := x.Index(i).Interface()
			valueType := reflect.ValueOf(value)

			wrongDataInput := valueType.Kind() == reflect.String && component.InputType != "String"

			if !wrongDataInput && slices.Contains(options, valueType.String()) {
				continue
			}

			log.Error(fmt.Sprintf("got string value for %s component %s", component.InputType, component.Source))
			return false, fmt.Errorf(fmt.Sprintf("wrong value for %s source", component.Source))
		}

		return true, nil

	default:
		log.Error("not supported value")
		return false, fmt.Errorf(fmt.Sprintf("wrong value for %s source", component.Source))
	}

}

func (s *Service) Validate(c []models.Component, payload map[string]any) (bool, error) {
	for _, c := range c {
		value, exists := payload[c.Source]
		if !exists {
			return false, fmt.Errorf("component %s is required, input type %s", c.Source, c.InputType)
		}

		valid, err := s.isValidValue(&c, value)

		if err != nil {
			return false, err
		}

		if !valid {
			return false, fmt.Errorf("wrong value for %s source", c.Source)
		}
	}

	return true, nil
}
