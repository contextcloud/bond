package tests

import (
	"bond/examples"
	"bond/pkg/parser"
	"bond/pkg/terra"
	"context"
	"testing"
)

func TestFactoryNew(t *testing.T) {
	testdata := []struct {
		Name string
	}{{
		Name: "buckets.hcl",
	}}

	ctx := context.Background()
	withBackendS3 := terra.WithBackendS3("contextcloud-bond-test-bucket", "us-east-1")
	withBaseDir := terra.WithBaseDir("./tmp")

	p := parser.NewParser()
	f, err := terra.NewFactory(ctx, withBackendS3, withBaseDir)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, d := range testdata {
		t.Run(d.Name, func(t *testing.T) {
			raw, err := examples.ReadFile(d.Name)
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
