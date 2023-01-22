package terra

import (
	"bond/pkg/parser"
	"context"
	"testing"

	"github.com/spf13/afero"
)

func TestFactoryNew(t *testing.T) {
	const exampleConfig = `
		resource "s3_bucket" "bar" {
			foo = "bar"
		}
	`

	p := parser.NewParser()

	cfg, err := p.Parse("main.hcl", []byte(exampleConfig))
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Run("creates new tf", func(t *testing.T) {
		ctx := context.Background()

		// todo
		fs := afero.NewOsFs()
		if err := fs.RemoveAll("tmp"); err != nil {
			t.Fatal(err)
			return
		}

		f, err := NewFactory(fs, "tmp")
		if err != nil {
			t.Fatal(err)
			return
		}

		if _, err := f.New(ctx, cfg); err != nil {
			t.Fatal(err)
			return
		}
	})
}
