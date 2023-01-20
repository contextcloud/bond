package terra

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type Dict map[string]interface{}

func (d Dict) String() string {
	return ""
}

type Config struct {
	Resources []*Resource `hcl:"resource,block"`
}

type Resource struct {
	Type    string `hcl:"type,label"`
	Name    string `hcl:"name,label"`
	Options Dict   `hcl:",remain"`
}

func Parse(filename string, data []byte) (*Config, error) {
	var cfg Config
	if err := hclsimple.Decode(filename, data, nil, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
