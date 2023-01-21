package terra

import (
	"context"
	"embed"
	"html/template"
	"os"
	"path"

	"github.com/spf13/afero"
)

//go:embed resources/*
var tmpls embed.FS

type Factory interface {
	New(ctx context.Context, cfg *Config) (Terraform, error)
}

type factory struct {
	fs      afero.Fs
	baseDir string
}

func (m *factory) createMain(tmp string, data interface{}) error {
	p := path.Join(tmp, "main.tf")
	f, err := m.fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	// create the terraform structure
	tpl, err := template.ParseFS(tmpls, "resources/aws.tf.tmpl")
	if err != nil {
		return err
	}
	if err := tpl.Execute(f, data); err != nil {
		return err
	}
	return nil
}
func (m *factory) writeModules(tmp string, modules []*Module) error {
	p := path.Join(tmp, "modules.tf")
	f, err := m.fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	data := Encode(modules)
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

func (m *factory) New(ctx context.Context, cfg *Config) (Terraform, error) {
	var modules = make([]*Module, len(cfg.Resources))
	for i, r := range cfg.Resources {
		// make the file.
		modules[i] = &Module{
			Name:    r.Name,
			Source:  "./" + r.Name,
			Options: r.Options,
		}
	}

	// create a temp dir.
	tmp, err := os.MkdirTemp(m.baseDir, "bond-")
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	if err := m.createMain(tmp, data); err != nil {
		return nil, err
	}

	if err := m.writeModules(tmp, modules); err != nil {

	}

	return &terraform{}, nil
}

func NewFactory(fs afero.Fs, baseDir string) (Factory, error) {
	if err := fs.MkdirAll(baseDir, 0755); err != nil {
		return nil, err
	}

	return &factory{
		fs:      fs,
		baseDir: baseDir,
	}, nil
}
