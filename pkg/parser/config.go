package parser

type Config struct {
	Resources []*Resource `hcl:"resource,block"`
}

type Resource struct {
	Type    string      `hcl:"type,label"`
	Name    string      `hcl:"name,label"`
	Options interface{} `hcl:",squash"`
}
