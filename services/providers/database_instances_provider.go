package providers

type DataBaseInstancesProvider struct{}

func (o *DataBaseInstancesProvider) GetComponentOptions() *ComponentTypes {
	return &ComponentTypes{
		"ComboBox": {"Prod", "Local", "staging"},
	}
}
