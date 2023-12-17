package providers

type DefaultOptionsProvider struct{}

func (o *DefaultOptionsProvider) GetComponentOptions() *ComponentTypes {
	return &ComponentTypes{
		"TextField":     {},
		"ComboBox":      {},
		"RadioButton":   {},
		"MultiComboBox": {},
	}
}
