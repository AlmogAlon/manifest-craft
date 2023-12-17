package services

type DataBaseInstancesProvider struct{}

func (o *DataBaseInstancesProvider) GetComponentOptions() *ComponentTypes {
	return &ComponentTypes{
		"ComboBox": {},
	}
}
