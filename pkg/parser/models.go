package parser

type Boundry struct {
	Id        string
	Env       map[string]string
	Providers []*Provider
	Resources []*Resource
}

type Provider struct {
	Name    string
	Options interface{}
}

type Resource struct {
	Type      string
	Name      string
	Options   interface{}
	DependsOn []string
}
