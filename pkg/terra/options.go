package terra

import (
	"bond/pkg/utils"
	"path/filepath"

	"github.com/spf13/afero"
)

type Options struct {
	Fs             afero.Fs
	Env            map[string]string
	ExecPath       string
	BaseDir        string
	BackendType    BackendType
	BackendOptions interface{}
	DryRun         bool
}

func NewOptions() *Options {
	osEnv := utils.EnvionmentVars()

	return &Options{
		Fs:          afero.NewOsFs(),
		Env:         osEnv,
		ExecPath:    "terraform",
		BaseDir:     ".",
		BackendType: BackendTypeLocal,
		DryRun:      false,
	}
}

type Option (func(*Options))

func WithBackendS3(bucket string, region string) Option {
	return func(o *Options) {
		o.BackendType = BackendTypeS3
		o.BackendOptions = &BackendS3{
			Bucket: bucket,
			Region: region,
		}
	}
}

func WithBaseDir(dir string) Option {
	return func(o *Options) {
		if dir == "" {
			return
		}

		o.BaseDir = dir
	}
}

func WithExecPath(path string) Option {
	return func(o *Options) {
		if path == "" {
			return
		}

		if path == "terraform" {
			o.ExecPath = path
			return
		}

		execPath, _ := filepath.Abs(path)
		o.ExecPath = execPath
	}
}

func WithDryRun(dryRun bool) Option {
	return func(o *Options) {
		o.DryRun = dryRun
	}
}

func AddEnv(env map[string]string) Option {
	return func(o *Options) {
		o.Env = utils.MergeMaps(o.Env, env)
	}
}
