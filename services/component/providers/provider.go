package providers

type ComponentTypes map[string][]string

type DefaultOptionsProvider struct{}

func (o *DefaultOptionsProvider) GetComponentOptions() *ComponentTypes {
	return &ComponentTypes{
		"TextField":     {},
		"ComboBox":      {},
		"RadioButton":   {},
		"MultiComboBox": {},
	}
}

type Provider interface {
	GetComponentOptions() *ComponentTypes
}
