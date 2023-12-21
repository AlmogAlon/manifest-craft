package services

import (
	"manifest-craft/services/component"
)

type Services struct {
	Component component.Interface
}

func Get() Services {
	return Services{
		Component: component.NewComponentService(),
	}
}
