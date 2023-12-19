package providers

type AccountsProvider struct{}

func (o *AccountsProvider) GetComponentOptions() (*ComponentTypes, error) {
	return &ComponentTypes{
		"MultiComboBox": {"alonalmog"},
	}, nil
}
