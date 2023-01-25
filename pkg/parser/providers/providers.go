package providers

type Factory func() interface{}
type Providers map[string]Factory

func NewProviders() Providers {
	return map[string]Factory{
		"aws": func() interface{} { return &AWSProvider{} },
	}
}
