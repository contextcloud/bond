package tests

import (
	"bond/config"
	"bond/examples"
	"bond/pkg/client"
	"bond/pkg/parser"
	"context"
	"testing"
)

func TestFactoryNew(t *testing.T) {
	testdata := []struct {
		Name string
	}{{
		Name: "cdn.hcl",
	}}

	ctx := context.Background()
	cfg := &config.Config{
		BaseDir: "./tmp",
		DryRun:  true,
	}

	p := parser.NewParser()
	f, err := client.NewFactory(ctx, cfg)
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

			c, err := f.New(ctx, cfg)
			if err != nil {
				t.Fatal(err)
				return
			}

			if err := c.Apply(ctx); err != nil {
				t.Fatal(err)
				return
			}
		})
	}
}
