package tests

import (
	"bond/examples"
	"bond/pkg/parser"
	"testing"
)

func TestParse(t *testing.T) {
	testdata := []struct {
		Name string
	}{{
		Name: "domains.hcl",
	}}

	p := parser.NewParser()

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

			t.Logf("cfg: %+v", cfg)
		})
	}

}
