package parser

import "github.com/hashicorp/hcl/v2"

var dependsOnSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "depends_on",
		},
	},
}

func DependsOn(body hcl.Body) ([]string, hcl.Body, error) {
	var diags hcl.Diagnostics
	partial, remains, partialDiags := body.PartialContent(dependsOnSchema)
	diags = append(diags, partialDiags...)

	var dependsOn []string

	for _, attr := range partial.Attributes {
		switch attr.Name {
		case "depends_on":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			l := v.AsValueSlice()
			dependsOn = make([]string, len(l))
			for i, v := range l {
				dependsOn[i] = v.AsString()
			}
		}
	}

	if diags.HasErrors() {
		return nil, nil, diags
	}
	return dependsOn, remains, nil
}

var providersSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "providers",
		},
	},
}

func Providers(body hcl.Body) (map[string]string, hcl.Body, error) {
	var diags hcl.Diagnostics
	partial, remains, partialDiags := body.PartialContent(providersSchema)
	diags = append(diags, partialDiags...)

	var providers map[string]string

	for _, attr := range partial.Attributes {
		switch attr.Name {
		case "providers":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			m := v.AsValueMap()
			providers = make(map[string]string, len(m))
			for k, v := range m {
				providers[k] = v.AsString()
			}
		}
	}

	if diags.HasErrors() {
		return nil, nil, diags
	}
	return providers, remains, nil
}

var aliasSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "alias",
		},
	},
}

func Alias(body hcl.Body) (string, hcl.Body, error) {
	var diags hcl.Diagnostics
	partial, remains, partialDiags := body.PartialContent(aliasSchema)
	diags = append(diags, partialDiags...)

	var alias string

	for _, attr := range partial.Attributes {
		switch attr.Name {
		case "alias":
			v, vDiags := attr.Expr.Value(nil)
			if vDiags.HasErrors() {
				diags = append(diags, vDiags...)
				continue
			}

			alias = v.AsString()
		}
	}

	if diags.HasErrors() {
		return "", nil, diags
	}
	return alias, remains, nil
}
