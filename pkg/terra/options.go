package terra

import (
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
}

func NewOptions() *Options {
	osEnv := EnvionmentVars()

	return &Options{
		Fs:          afero.NewOsFs(),
		Env:         osEnv,
		ExecPath:    "terraform",
		BaseDir:     ".",
		BackendType: BackendTypeLocal,
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
		o.BaseDir = dir
	}
}

func WithExecPath(path string) Option {
	return func(o *Options) {
		execPath, _ := filepath.Abs(path)
		o.ExecPath = execPath
	}
}

func AddEnv(env map[string]string) Option {
	return func(o *Options) {
		o.Env = MergeMaps(o.Env, env)
	}
}
