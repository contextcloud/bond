package terra

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
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
	fs       afero.Fs
	execPath string
	env      map[string]string
	baseDir  string
	backend  *Backend
}

func (m *factory) open(p string) (afero.File, error) {
	f, err := m.fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
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
		if err := modules.CopyModule(m.fs, tmp, k); err != nil {
			return err
		}
	}

	return nil
}

func (m *factory) pushState(ctx context.Context, tf *tfexec.Terraform, id string) error {
	stateFile := path.Join(m.baseDir, "states", id+".tfstate")
	if _, err := m.fs.Stat(stateFile); errors.Is(err, os.ErrNotExist) {
		return nil
	} else if err != nil {
		return err
	}

	// do we need to push existing state?
	path := path.Join("../states/", id+".tfstate")
	if err := tf.StatePush(ctx, path); err != nil {
		return err
	}
	if err := m.fs.Remove(stateFile); err != nil {
		// not sure what to do here!.
	}
	return nil
}

func (m *factory) New(ctx context.Context, cfg *parser.Boundry) (Terraform, error) {
	// create a temp dir.
	tmp, err := os.MkdirTemp(m.baseDir, "boundries-")
	if err != nil {
		return nil, err
	}

	env := MergeMaps(m.env, cfg.Env)

	tf, err := tfexec.NewTerraform(tmp, m.execPath)
	if err != nil {
		return nil, err
	}
	if err := tf.SetEnv(env); err != nil {
		return nil, err
	}

	// write the provider again this time with remote backend
	if err := m.createProviders(tmp, cfg, m.backend.Type, m.backend.Options); err != nil {
		return nil, err
	}
	if err := m.createMain(tmp, cfg.Resources); err != nil {
		return nil, err
	}

	if err := tf.Init(ctx, tfexec.Upgrade(true)); err != nil {
		return nil, err
	}

	if m.backend.Type != BackendTypeLocal {
		if err := m.pushState(ctx, tf, cfg.Id); err != nil {
			return nil, err
		}
	}

	return &terraform{
		tf: tf,
	}, nil
}

func NewFactory(ctx context.Context, fs afero.Fs, env map[string]string, baseDir string, backend *Backend) (Factory, error) {
	statesDir := path.Join(baseDir, "states")
	if err := fs.MkdirAll(statesDir, 0755); err != nil {
		return nil, err
	}

	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.3.7")),
	}

	execPath, err := installer.Install(ctx)
	if err != nil {
		return nil, err
	}

	b := backend
	if b == nil {
		b = &Backend{
			Type: BackendTypeLocal,
		}
	}

	return &factory{
		fs:       fs,
		execPath: execPath,
		env:      env,
		baseDir:  baseDir,
		backend:  b,
	}, nil
}
