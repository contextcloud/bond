package examples

import "embed"

//go:embed **
var data embed.FS

func ReadFile(name string) ([]byte, error) {
	return data.ReadFile(name)
}
