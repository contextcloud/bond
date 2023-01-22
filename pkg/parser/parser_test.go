package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	const exampleConfig = `
		resource "s3_bucket" "bar" {
			foo = "bar"
			foo2 = "bar2"
		}
	`

	p := NewParser()
	cfg, err := p.Parse("main.hcl", []byte(exampleConfig))
	if err != nil {
		t.Fatal(err)
		return
	}

	if len(cfg.Resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(cfg.Resources))
		return
	}
	if cfg.Resources[0].Type != "s3_bucket" {
		t.Fatalf("expected resource type 'foo', got '%s'", cfg.Resources[0].Type)
		return
	}
	if cfg.Resources[0].Name != "bar" {
		t.Fatalf("expected resource name 'bar', got '%s'", cfg.Resources[0].Name)
		return
	}
}
