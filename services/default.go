package services

type DefaultOptionsProvider struct{}

func (o *DefaultOptionsProvider) GetComponentOptions() *ComponentTypes {
	return &ComponentTypes{
		"TextField":     {},
		"ComboBox":      {"Read Only", "Read write", "Full Admin"},
		"RadioButton":   {},
		"MultiComboBox": {},
	}
}
