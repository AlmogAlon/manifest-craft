package providers

type DefaultOptionsProvider struct{}

func (o *DefaultOptionsProvider) GetComponentOptions() (*ComponentTypes, error) {
	return &ComponentTypes{
		"ComboBox":      {},
		"TextField":     {},
		"RadioButton":   {},
		"MultiComboBox": {},
	}, nil
}
