package parser

import (
	"bond/pkg/parser/providers"
	"bond/pkg/parser/resources"
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Parser interface {
	Parse(filename string, data []byte) (*Boundry, error)
}

type parser struct {
	providers providers.Providers
	resources resources.Resources
}

func (p *parser) load(filename string, data []byte) (*hcl.File, error) {
	parser := hclparse.NewParser()

	if strings.HasSuffix(filename, ".json") {
		file, fileDiags := parser.ParseJSON(data, filename)
		if fileDiags.HasErrors() {
			return nil, fileDiags
		}
		return file, nil
	}

	file, fileDiags := parser.ParseHCL(data, filename)
	if fileDiags.HasErrors() {
		return nil, fileDiags
	}
	return file, nil
}

func (p *parser) getProvider(name string, body hcl.Body) (*Provider, error) {
	factory, ok := p.providers[name]
	if !ok {
		return nil, fmt.Errorf("unknown provider %q", name)
	}

	opts, err := factory(body)
	if err != nil {
		return nil, err
	}

	return &Provider{
		Name:    name,
		Options: opts,
	}, nil
}

func (p *parser) getResource(typeName string, name string, body hcl.Body) (*Resource, error) {
	factory, ok := p.resources[typeName]
	if !ok {
		return nil, fmt.Errorf("unknown resource type %q", typeName)
	}

	dependsOn, remains, err := DependsOn(body)
	if err != nil {
		return nil, err
	}

	opts, err := factory(remains)
	if err != nil {
		return nil, err
	}

	return &Resource{
		Type:      typeName,
		Name:      name,
		Options:   opts,
		DependsOn: dependsOn,
	}, nil
}

func (p *parser) Parse(filename string, data []byte) (*Boundry, error) {
	// load the file.
	file, err := p.load(filename, data)
	if err != nil {
		return nil, err
	}

	cfg := &Boundry{
		Env: map[string]string{},
	}

	// create the eval context

	var diags hcl.Diagnostics
	content, _, contentDiags := file.Body.PartialContent(rootSchema)
	diags = append(diags, contentDiags...)

	for _, attr := range content.Attributes {
		switch attr.Name {
		case "id":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}
			cfg.Id = v.AsString()
		}
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "env":
			attrs, attrsDiags := block.Body.JustAttributes()
			if attrsDiags.HasErrors() {
				diags = append(diags, attrsDiags...)
				continue
			}

			for k, attr := range attrs {
				v, vDiags := attr.Expr.Value(nil)
				if vDiags.HasErrors() {
					diags = append(diags, vDiags...)
					continue
				}
				cfg.Env[k] = v.AsString()
			}

		case "provider":
			name := block.Labels[0]

			provider, err := p.getProvider(name, block.Body)
			if err != nil {
				diags = diags.Append(&hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "Invalid provider",
					Detail:   fmt.Sprintf("The provider %q is not valid. %s", name, err.Error()),
					Subject:  &block.DefRange,
				})
				continue
			}
			cfg.Providers = append(cfg.Providers, provider)
		case "resource":
			typeName := block.Labels[0]
			name := block.Labels[1]

			resource, err := p.getResource(typeName, name, block.Body)
			if err != nil {
				diags = diags.Append(&hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "Invalid resource type",
					Detail:   fmt.Sprintf("The resource type %q is not valid. %s", typeName, err.Error()),
					Subject:  &block.DefRange,
				})
				continue
			}
			cfg.Resources = append(cfg.Resources, resource)
		}
	}

	if diags.HasErrors() {
		return nil, diags
	}
	return cfg, nil
}

func NewParser() Parser {
	return &parser{
		providers: providers.NewProviders(),
		resources: resources.NewResources(),
	}
}
