package terra

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-exec/tfexec"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"

	"bond/modules"
	"bond/pkg/parser"
)

type Factory interface {
	New(ctx context.Context, cfg *parser.Boundry) (Terraform, error)
}

type factory struct {
	*Options
}

func (m *factory) ensure(dir string) (string, error) {
	// check if the dir exists.
	p := path.Join(m.BaseDir, dir)

	fi, err := m.Fs.Stat(p)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	if err == nil && fi.IsDir() {
		// clean up the dir
		modules := path.Join(p, "modules")
		if err := m.Fs.RemoveAll(modules); err != nil {
			return "", err
		}
		if err := m.Fs.RemoveAll(path.Join(p, "main.tf")); err != nil {
			return "", err
		}
		if err := m.Fs.RemoveAll(path.Join(p, "providers.tf")); err != nil {
			return "", err
		}
		return p, nil
	}
	if err == nil && !fi.IsDir() {
		if err := m.Fs.Remove(p); err != nil {
			return "", err
		}
	}

	// create a temp dir.
	if err := os.MkdirAll(p, 0755); err != nil {
		return "", err
	}

	return p, nil
}

func (m *factory) open(p string) (afero.File, error) {
	f, err := m.Fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	if err := f.Truncate(0); err != nil {
		return nil, err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}
	return f, nil
}

func (m *factory) createProviders(tmp string, boundry *parser.Boundry, backendType BackendType, backendOptions interface{}) error {
	p := path.Join(tmp, "providers.tf")
	f, err := m.open(p)
	if err != nil {
		return err
	}
	defer f.Close()

	terraformBlock := gohcl.EncodeAsBlock(NewTerraformBlock(), "terraform")

	switch backendType {
	case BackendTypeS3:
		backendS3Block := gohcl.EncodeAsBlock(backendOptions, "backend")
		backendS3Block.Body().SetAttributeValue("key", cty.StringVal(boundry.Id+".tfstate"))
		backendS3Block.SetLabels([]string{"s3"})
		terraformBlock.Body().AppendBlock(backendS3Block)
		break
	default:
		backend := &BackendLocal{
			Path: path.Join("../states/", boundry.Id+".tfstate"),
		}
		backendLocalBlock := gohcl.EncodeAsBlock(backend, "backend")
		backendLocalBlock.SetLabels([]string{"local"})
		terraformBlock.Body().AppendBlock(backendLocalBlock)
		break
	}

	writer := hclwrite.NewEmptyFile()
	writer.Body().AppendBlock(terraformBlock)

	for _, m := range boundry.Providers {
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
	f, err := m.open(p)
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
		if err := modules.CopyModule(m.Fs, tmp, k); err != nil {
			return err
		}
	}

	return nil
}

func (m *factory) New(ctx context.Context, cfg *parser.Boundry) (Terraform, error) {
	workingDir, err := m.ensure(cfg.Id)
	if err != nil {
		return nil, err
	}

	env := MergeMaps(m.Env, cfg.Env)

	tf, err := tfexec.NewTerraform(workingDir, m.ExecPath)
	if err != nil {
		return nil, err
	}
	if err := tf.SetEnv(env); err != nil {
		return nil, err
	}

	// write the provider again this time with remote backend
	if err := m.createProviders(workingDir, cfg, m.BackendType, m.BackendOptions); err != nil {
		return nil, err
	}
	if err := m.createMain(workingDir, cfg.Resources); err != nil {
		return nil, err
	}

	if err := tf.Init(ctx, tfexec.Upgrade(true)); err != nil {
		return nil, err
	}

	return &terraform{
		tf: tf,
	}, nil
}

func NewFactory(ctx context.Context, opts ...Option) (Factory, error) {
	o := NewOptions()
	for _, opt := range opts {
		opt(o)
	}

	return &factory{
		Options: o,
	}, nil
}
