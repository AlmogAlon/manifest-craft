package providers

type AccountsProvider struct{}

func (o *AccountsProvider) GetComponentOptions() *ComponentTypes {
	return &ComponentTypes{
		"MultiComboBox": {"alonalmog"},
	}
}
