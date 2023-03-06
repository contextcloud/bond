package terra

import (
	"bond/modules"
	"bond/pkg/parser"
	"fmt"
	"os"
	"path"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/spf13/afero"
	"github.com/zclconf/go-cty/cty"
)

func writeDependsOn(block *hclwrite.Block, dependsOn []string) {
	if len(dependsOn) == 0 {
		return
	}

	toks := hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenOBrack,
			Bytes: []byte("["),
		},
	}

	for i, d := range dependsOn {
		if i > 0 {
			toks = append(toks, &hclwrite.Token{
				Type:  hclsyntax.TokenComma,
				Bytes: []byte{','},
			})
		}
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(fmt.Sprintf("module.%s", d)),
		})
	}

	toks = append(toks, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrack,
		Bytes: []byte("]"),
	})

	block.Body().SetAttributeRaw("depends_on", toks)
}

func writeProviders(block *hclwrite.Block, providers map[string]string) {
	if len(providers) == 0 {
		return
	}

	toks := hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenOBrace,
			Bytes: []byte("{"),
		},
	}

	count := 0
	for k, v := range providers {
		if count > 0 {
			toks = append(toks, &hclwrite.Token{
				Type:  hclsyntax.TokenComma,
				Bytes: []byte{','},
			})
		}
		count++

		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(k),
		})
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenEqual,
			Bytes: []byte("="),
		})
		toks = append(toks, &hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(v),
		})
	}

	toks = append(toks, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrace,
		Bytes: []byte("}"),
	})

	block.Body().SetAttributeRaw("providers", toks)
}

func ensure(fs afero.Fs, baseDir string, dir string) (string, error) {
	// check if the dir exists.
	p := path.Join(baseDir, dir)

	fi, err := fs.Stat(p)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	if err == nil && fi.IsDir() {
		for _, name := range paths {
			if err := fs.RemoveAll(path.Join(p, name)); err != nil {
				return "", err
			}
		}
		return p, nil
	}
	if err == nil && !fi.IsDir() {
		if err := fs.Remove(p); err != nil {
			return "", err
		}
	}

	// create a temp dir.
	if err := fs.MkdirAll(p, 0755); err != nil {
		return "", err
	}

	return p, nil
}

func open(fs afero.Fs, p string) (afero.File, error) {
	f, err := fs.OpenFile(p, os.O_RDWR|os.O_CREATE, 0755)
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

func createProviders(fs afero.Fs, dir string, boundry *parser.Boundry, backendType BackendType, backendOptions interface{}) error {
	p := path.Join(dir, "providers.tf")
	f, err := open(fs, p)
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
		if len(m.Alias) > 0 {
			block.Body().SetAttributeValue("alias", cty.StringVal(m.Alias))
		}
		block.SetLabels([]string{m.Name})
		writer.Body().AppendBlock(block)
	}
	if _, err := writer.WriteTo(f); err != nil {
		return err
	}

	return nil
}
func createMain(fs afero.Fs, dir string, resources []*parser.Resource) error {
	p := path.Join(dir, "main.tf")
	f, err := open(fs, p)
	if err != nil {
		return err
	}
	defer f.Close()

	mods := map[string]bool{}

	writer := hclwrite.NewEmptyFile()
	for _, r := range resources {
		mods[r.Type] = true

		// source
		source := fmt.Sprintf("./modules/%s", r.Type)

		block := gohcl.EncodeAsBlock(r.Options, "module")
		block.Body().SetAttributeValue("source", cty.StringVal(source))
		block.SetLabels([]string{r.Name})

		writeDependsOn(block, r.DependsOn)
		writeProviders(block, r.Providers)

		writer.Body().AppendBlock(block)
	}
	if _, err := writer.WriteTo(f); err != nil {
		return err
	}

	// copy the modules.
	for k := range mods {
		if err := modules.CopyModule(fs, dir, k); err != nil {
			return err
		}
	}

	return nil
}
func createOutputs(fs afero.Fs, dir string, resources []*parser.Resource) error {
	p := path.Join(dir, "outputs.tf")
	f, err := open(fs, p)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := hclwrite.NewEmptyFile()
	for _, m := range resources {
		toks := hclwrite.Tokens{
			&hclwrite.Token{
				Type:  hclsyntax.TokenIdent,
				Bytes: []byte(fmt.Sprintf("module.%s", m.Name)),
			},
		}

		block := hclwrite.NewBlock("output", []string{m.Name})
		block.Body().SetAttributeRaw("value", toks)
		writer.Body().AppendBlock(block)
	}
	if _, err := writer.WriteTo(f); err != nil {
		return err
	}

	return nil
}
