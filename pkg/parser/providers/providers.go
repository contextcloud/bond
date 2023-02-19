package providers

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

type Provider interface {
}
type Factory func(hcl.Body) (Provider, error)
type Providers map[string]Factory

func Basic(body hcl.Body, opts Provider) (Provider, error) {
	if err := gohcl.DecodeBody(body, nil, opts); err != nil {
		return nil, err
	}
	return opts, nil
}

func NewProviders() Providers {
	return map[string]Factory{
		"aws": func(b hcl.Body) (Provider, error) { return Basic(b, &AWSProvider{}) },
	}
}
