package providers

type DefaultOptionsProvider struct{}

func (o *DefaultOptionsProvider) GetComponentOptions() (*ComponentTypes, error) {
	return &ComponentTypes{
		"TextField":     {},
		"ComboBox":      {},
		"RadioButton":   {},
		"MultiComboBox": {},
	}, nil
}
