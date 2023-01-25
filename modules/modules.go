package modules

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"

	"github.com/spf13/afero"
)

//go:embed **
var modules embed.FS

func CopyModule(afs afero.Fs, tmp string, name string) error {
	return fs.WalkDir(modules, name, func(n string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		tmpl, err := modules.Open(n)
		if err != nil {
			return err
		}

		// remove prefix
		o := path.Join(tmp, "modules", n)

		dir := path.Dir(o)
		if err := afs.MkdirAll(dir, 0755); err != nil {
			return err
		}

		f, err := afs.OpenFile(o, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}

		if _, err := io.Copy(f, tmpl); err != nil {
			return err
		}

		return nil
	})
}
