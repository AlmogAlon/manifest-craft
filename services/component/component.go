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

func NewComponentService() Interface {
	return &Service{
		Providers: providers.InitProviders(),
	}
}

type Interface interface {
	Validate(c []models.Component, payload map[string]any) (bool, error)
	GetComponentOptions(component *models.Component) ([]string, error)
}

func (s *Service) isCorrectType(component *models.Component, k reflect.Kind) bool {
	switch k {
	case reflect.Bool:
		return component.ComponentType == "RadioButton" && component.InputType == "Boolean"
	case reflect.String:
		supportedStringTypes := []string{"TextField", "ComboBox", "MultiComboBox"}
		return component.InputType == "String" && slices.Contains(supportedStringTypes, component.ComponentType)
	case reflect.Slice:
		return component.InputType == "String" && component.ComponentType == "MultiComboBox"

	default:
		return false
	}
}

func (s *Service) isValidValue(component *models.Component, v interface{}) (bool, error) {
	x := reflect.ValueOf(v)

	vKind := x.Kind()

	if !s.isCorrectType(component, vKind) {
		log.Error(fmt.Sprintf("got incorrect value for %s component %s", component.InputType, component.Source))
		return false, fmt.Errorf(fmt.Sprintf("wrong value for %s source", component.Source))
	}
	switch vKind {
	case reflect.Bool:
		return true, nil

	case reflect.String:
		if component.ComponentType == "TextField" {
			return true, nil
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

			wrongDataInput := !s.isCorrectType(component, valueType.Kind())

			if !wrongDataInput && slices.Contains(options, valueType.String()) {
				continue
			}

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
