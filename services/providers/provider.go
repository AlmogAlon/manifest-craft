package providers

type ComponentTypes map[string][]string

type Provider interface {
	GetComponentOptions() *ComponentTypes
}
