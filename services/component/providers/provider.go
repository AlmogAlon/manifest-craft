package providers

type ComponentTypes map[string][]string

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
