package parser

import "github.com/hashicorp/hcl/v2"

var rootSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
	},
}
