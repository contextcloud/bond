package terra

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"

	"bond/pkg/parser"
)

//go:embed resources/*
var tmpls embed.FS

type Factory interface {
	New(ctx context.Context, cfg *parser.Config) (Terraform, error)
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
func (m *factory) copyFiles(tmp string, source string) error {
	prefix := "resources"

	p := path.Join(prefix, source)
	return fs.WalkDir(tmpls, p, func(n string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// do nothing
		if d.IsDir() {
			return nil
		}

		tmpl, err := tmpls.Open(n)
		if err != nil {
			return err
		}

		// remove prefix
		o := strings.TrimPrefix(n, prefix)
		o = path.Join(tmp, o)

		dir := path.Dir(o)
		if err := m.fs.MkdirAll(dir, 0755); err != nil {
			return err
		}

		f, err := m.fs.OpenFile(o, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}

		if _, err := io.Copy(f, tmpl); err != nil {
			return err
		}

		return nil
	})
}
func (m *factory) writeModules(tmp string, provider string, modules []*parser.Resource) error {
	p := path.Join(tmp, "modules.tf")
	f, err := m.fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	sources := map[string]bool{}

	writer := hclwrite.NewEmptyFile()
	for _, m := range modules {
		// source
		source := fmt.Sprintf("./modules/%s/%s", provider, m.Type)
		sources[source] = true

		block := gohcl.EncodeAsBlock(m.Options, "module")
		block.Body().SetAttributeValue("source", cty.StringVal(source))
		block.SetLabels([]string{m.Name})
		writer.Body().AppendBlock(block)
	}
	if _, err := writer.WriteTo(f); err != nil {
		return err
	}

	// copy the modules.
	for k := range sources {
		if err := m.copyFiles(tmp, k); err != nil {
			return err
		}
	}

	return nil
}

func (m *factory) New(ctx context.Context, cfg *parser.Config) (Terraform, error) {
	// create a temp dir.
	tmp, err := os.MkdirTemp(m.baseDir, "bond-")
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	if err := m.createMain(tmp, data); err != nil {
		return nil, err
	}

	if err := m.writeModules(tmp, "aws", cfg.Resources); err != nil {
		return nil, err
	}

	return NewTerraform(ctx, tmp)
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
