package terra

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Config struct {
	Resources []*Resource `hcl:"resource,block"`
}

type Resource struct {
	Type    string                 `hcl:"type,label"`
	Name    string                 `hcl:"name,label"`
	Options map[string]interface{} `hcl:",remain"`
}

type Module struct {
	Name    string      `hcl:"name,label"`
	Source  string      `hcl:"source"`
	Options interface{} `hcl:",squash"`
}

func Parse(filename string, data []byte) (*Config, error) {
	var cfg Config
	if err := hclsimple.Decode(filename, data, nil, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func Encode(val interface{}) []byte {
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(val, f.Body())
	return f.Bytes()
}
