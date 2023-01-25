package parser

type Config struct {
	Env       map[string]string `hcl:"env,block"`
	Providers []*Provider       `hcl:"provider,block"`
	Resources []*Resource       `hcl:"resource,block"`
}

type Provider struct {
	Name    string      `hcl:"name,label"`
	Options interface{} `hcl:",squash"`
}

type Resource struct {
	Type    string      `hcl:"type,label"`
	Name    string      `hcl:"name,label"`
	Options interface{} `hcl:",squash"`
}
