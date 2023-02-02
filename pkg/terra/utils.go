package terra

import (
	"os"
	"strings"
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
