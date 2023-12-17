package services

import (
	"manifest-craft/services/component"
)

type Services struct {
	Component *component.Service
}

func Get() *Services {
	return &Services{
		Component: component.NewComponentService(),
	}
}
