package providers

type DataBaseInstancesProvider struct{}

func (o *DataBaseInstancesProvider) GetComponentOptions() (*ComponentTypes, error) {
	return &ComponentTypes{
		"ComboBox": {"Prod", "Local", "staging"},
	}, nil
}
