package tests

import (
	"bond/pkg/parser"
	"bond/tests/data"
	"testing"
)

func TestParse(t *testing.T) {
	testdata := []struct {
		Name string
	}{{
		Name: "buckets.hcl",
	}}

	p := parser.NewParser()

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

			if len(cfg.Env) != 3 {
				t.Fatalf("expected 3 environment variables, got %d", len(cfg.Env))
				return
			}

			if len(cfg.Providers) != 1 {
				t.Fatalf("expected 1 provider, got %d", len(cfg.Providers))
				return
			}
			if cfg.Providers[0].Name != "aws" {
				t.Fatalf("expected provider name 'aws', got '%s'", cfg.Providers[0].Name)
				return
			}

			if len(cfg.Resources) != 2 {
				t.Fatalf("expected 2 resource, got %d", len(cfg.Resources))
				return
			}
			if cfg.Resources[0].Type != "aws_s3_bucket" {
				t.Fatalf("expected resource type 'foo', got '%s'", cfg.Resources[0].Type)
				return
			}
			if cfg.Resources[0].Name != "test-bucket" {
				t.Fatalf("expected resource name 'test-bucket', got '%s'", cfg.Resources[0].Name)
				return
			}
		})
	}

}
