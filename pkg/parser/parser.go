package parser

import (
	"bond/pkg/parser/plugins"
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Parser interface {
	Parse(filename string, data []byte) (*Config, error)
}

type parser struct {
	plugins plugins.Plugins
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

func (p *parser) getDecoder(typeName string) (plugins.Decoder, error) {
	plugin, ok := p.plugins[typeName]
	if !ok {
		return nil, fmt.Errorf("unknown resource type %q", typeName)
	}
	return plugin, nil
}

func (p *parser) Parse(filename string, data []byte) (*Config, error) {
	// load the file.
	file, err := p.load(filename, data)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	var diags hcl.Diagnostics
	content, _, contentDiags := file.Body.PartialContent(rootSchema)
	diags = append(diags, contentDiags...)

	for _, block := range content.Blocks {
		switch block.Type {
		case "resource":
			typeName := block.Labels[0]
			name := block.Labels[1]

			// Totally-hypothetical plugin manager (not part of HCL)
			decoder, err := p.getDecoder(typeName)
			if err != nil {
				diags = diags.Append(&hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "Invalid resource type",
					Detail:   fmt.Sprintf("The resource type %q is not valid.", typeName),
					Subject:  &block.DefRange,
				})
				continue
			}

			opts, err := decoder(block.Body)
			if err != nil {
				diags = diags.Append(&hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "Invalid resource options",
					Detail:   fmt.Sprintf("The resource options are not valid: %s.", err),
					Subject:  &block.DefRange,
				})
			}

			r := &Resource{
				Type:    typeName,
				Name:    name,
				Options: opts,
			}
			cfg.Resources = append(cfg.Resources, r)
		}
	}

	if diags.HasErrors() {
		return nil, diags
	}
	return cfg, nil
}

func NewParser() Parser {
	return &parser{
		plugins: plugins.NewPlugins(),
	}
}
