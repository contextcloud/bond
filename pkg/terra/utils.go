package terra

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func EnvionmentVars() map[string]string {
	m := make(map[string]string)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			m[e[:i]] = e[i+1:]
		}
	}
	return m
}

func MergeMaps(maps ...map[string]string) map[string]string {
	out := map[string]string{}
	for _, m := range maps {
		if m == nil {
			continue
		}
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

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
