package tests

import (
	"bond/pkg/parser"
	"bond/pkg/terra"
	"bond/tests/data"
	"context"
	"testing"

	"github.com/spf13/afero"
)

func TestFactoryNew(t *testing.T) {
	testdata := []struct {
		Name string
	}{{
		Name: "buckets.hcl",
	}}

	ctx := context.Background()

	backend := &terra.Backend{
		Type: terra.BackendTypeS3,
		Options: terra.BackendS3{
			Bucket: "contextcloud-bond-test-bucket",
			Region: "us-east-1",
		},
	}

	p := parser.NewParser()
	fs := afero.NewOsFs()
	f, err := terra.NewFactory(ctx, fs, "tmp", backend)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, d := range testdata {
		t.Run(d.Name, func(t *testing.T) {
			raw, err := data.ReadFile(d.Name)
			if err != nil {
				t.Fatal(err)
				return
			}

			cfg, err := p.Parse("main.hcl", raw)
			if err != nil {
				t.Fatal(err)
				return
			}

			ctx := context.Background()

			tf, err := f.New(ctx, cfg)
			if err != nil {
				t.Fatal(err)
				return
			}

			if _, err := tf.Plan(ctx); err != nil {
				t.Fatal(err)
				return
			}

			if err := tf.Apply(ctx); err != nil {
				t.Fatal(err)
				return
			}
		})
	}
}
