package providers

type DataBaseRulesProvider struct{}

func (o *DataBaseRulesProvider) GetComponentOptions() *ComponentTypes {
	return &ComponentTypes{
		"ComboBox": {"Full Admin", "User"},
	}
}
