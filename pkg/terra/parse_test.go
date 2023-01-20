package terra

import (
	"testing"
)

func TestParse(t *testing.T) {
	const exampleConfig = `
		resource "foo" "bar" {
			foo = "bar"
		}
	`

	cfg, err := Parse("main.hcl", []byte(exampleConfig))
	if err != nil {
		t.Fatal(err)
		return
	}

	if len(cfg.Resources) != 1 {
		t.Fatalf("expected 1 resource, got %d", len(cfg.Resources))
		return
	}
	if cfg.Resources[0].Type != "foo" {
		t.Fatalf("expected resource type 'foo', got '%s'", cfg.Resources[0].Type)
		return
	}
	if cfg.Resources[0].Name != "bar" {
		t.Fatalf("expected resource name 'bar', got '%s'", cfg.Resources[0].Name)
		return
	}
	// if cfg.Resources[0].Options["foo"] != "bar" {
	// 	t.Fatalf("expected resource option 'foo' to be 'bar', got '%s'", cfg.Resources[0].Options["foo"])
	// 	return
	// }
}
