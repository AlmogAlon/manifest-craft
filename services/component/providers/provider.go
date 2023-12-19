package providers

type ComponentTypes map[string][]string

type DefaultOptionsProvider struct{}

func (o *DefaultOptionsProvider) GetComponentOptions() (*ComponentTypes, error) {
	return &ComponentTypes{
		"TextField":     {},
		"ComboBox":      {},
		"RadioButton":   {},
		"MultiComboBox": {},
	}, nil
}

type Provider interface {
	GetComponentOptions() (*ComponentTypes, error)
}

func InitProviders() map[string]Provider {
	return map[string]Provider{
		"databaseInstances": &DataBaseInstancesProvider{},
		"databaseRoles":     &DataBaseRulesProvider{},
		"users":             &UsersProvider{},
		"accounts":          &AccountsProvider{},

		"countries": &CountriesProvider{},
	}
}
