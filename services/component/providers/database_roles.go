package providers

type DataBaseRulesProvider struct{}

func (o *DataBaseRulesProvider) GetComponentOptions() (*ComponentTypes, error) {
	return &ComponentTypes{
		"ComboBox": {"Full Admin", "User"},
	}, nil
}
