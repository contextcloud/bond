package parser

import "github.com/hashicorp/hcl/v2"

var rootSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name:     "id",
			Required: true,
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "env",
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
	},
}
