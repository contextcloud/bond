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

func (m *factory) create(tmp string, filename string, templateName string, data interface{}) error {
	p := path.Join(tmp, filename)
	f, err := m.fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	// create the terraform structure
	tpl, err := template.ParseFS(tmpls, templateName)
	if err != nil {
		return err
	}
	if err := tpl.Execute(f, data); err != nil {
		return err
	}
	return nil
}

func (m *factory) New(ctx context.Context, cfg *Config) (Terraform, error) {
	// create a temp dir.
	tmp, err := os.MkdirTemp(m.baseDir, "bond-")
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	if err := m.create(tmp, "main.tf", "resources/aws.tf.tmpl", data); err != nil {
		return nil, err
	}

	for _, r := range cfg.Resources {
		if err := m.create(tmp, r.Name+".tf", "resources/resource.tf.tmpl", r); err != nil {
			return nil, err
		}
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
