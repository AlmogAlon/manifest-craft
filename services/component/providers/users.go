package providers

type UsersProvider struct{}

func (o *UsersProvider) GetComponentOptions() (*ComponentTypes, error) {
	return &ComponentTypes{
		"ComboBox": {"test@test.com"},
	}, nil
}
