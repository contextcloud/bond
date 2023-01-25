package terra

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"

	"bond/modules"
	"bond/pkg/parser"
)

type Factory interface {
	New(ctx context.Context, cfg *parser.Config) (Terraform, error)
}

type factory struct {
	fs      afero.Fs
	baseDir string
}

func (m *factory) createProviders(tmp string, providers []*parser.Provider) error {
	p := path.Join(tmp, "providers.tf")
	f, err := m.fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := hclwrite.NewEmptyFile()

	block := gohcl.EncodeAsBlock(NewTerraformBlock(), "terraform")
	writer.Body().AppendBlock(block)

	for _, m := range providers {
		// source
		block := gohcl.EncodeAsBlock(m.Options, "provider")
		block.SetLabels([]string{m.Name})
		writer.Body().AppendBlock(block)
	}
	if _, err := writer.WriteTo(f); err != nil {
		return err
	}

	return nil
}
func (m *factory) createMain(tmp string, resources []*parser.Resource) error {
	p := path.Join(tmp, "main.tf")
	f, err := m.fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	mods := map[string]bool{}

	writer := hclwrite.NewEmptyFile()
	for _, m := range resources {
		mods[m.Type] = true

		// source
		source := fmt.Sprintf("./modules/%s", m.Type)

		block := gohcl.EncodeAsBlock(m.Options, "module")
		block.Body().SetAttributeValue("source", cty.StringVal(source))
		block.SetLabels([]string{m.Name})
		writer.Body().AppendBlock(block)
	}
	if _, err := writer.WriteTo(f); err != nil {
		return err
	}

	// copy the modules.
	for k := range mods {
		if err := modules.CopyModule(m.fs, tmp, k); err != nil {
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

	if err := m.createProviders(tmp, cfg.Providers); err != nil {
		return nil, err
	}

	if err := m.createMain(tmp, cfg.Resources); err != nil {
		return nil, err
	}

	return NewTerraform(ctx, tmp, cfg.Env)
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
