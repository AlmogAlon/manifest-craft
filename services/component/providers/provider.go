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

func InitProviders() map[string]Provider {
	return map[string]Provider{
		"databaseInstances": &DataBaseInstancesProvider{},
		"databaseRoles":     &DataBaseRulesProvider{},
	}
}
