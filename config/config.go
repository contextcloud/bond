package config

import (
	"bond/pkg/terra"
	"context"

	"github.com/spf13/viper"
)

type Config struct {
	BaseDir     string `mapstructure:"base_dir"`
	ExecPath    string `mapstructure:"exec_path"`
	AwsS3Bucket string `mapstructure:"aws_s3_bucket"`
	AwsS3Region string `mapstructure:"aws_s3_region"`
}

func NewConfig() (*Config, error) {
	c := &Config{
		BaseDir:     ".",
		ExecPath:    "./.bin/terraform",
		AwsS3Region: "us-east-1",
	}

	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

func NewTerraFactory(ctx context.Context, cfg *Config) (terra.Factory, error) {
	opts := []terra.Option{
		terra.WithBaseDir(cfg.BaseDir),
		terra.WithExecPath(cfg.ExecPath),
	}

	if len(cfg.AwsS3Bucket) > 0 && len(cfg.AwsS3Region) > 0 {
		opts = append(opts, terra.WithBackendS3(cfg.AwsS3Bucket, cfg.AwsS3Region))
	}

	terraFactory, err := terra.NewFactory(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return terraFactory, nil
}
