package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BaseDir     string `mapstructure:"base_dir"`
	ExecPath    string `mapstructure:"exec_path"`
	AwsS3Bucket string `mapstructure:"aws_s3_bucket"`
	AwsS3Region string `mapstructure:"aws_s3_region"`
	DryRun      bool   `mapstructure:"dry_run"`
}

func NewConfig() (*Config, error) {
	c := &Config{
		BaseDir:     ".",
		ExecPath:    "./.bin/terraform",
		AwsS3Region: "us-east-1",
		DryRun:      false,
	}

	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}
