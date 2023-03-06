package terra

import (
	"context"

	"bond/config"
	"bond/pkg/parser"
)

var paths = []string{
	"modules",
	"main.tf",
	"providers.tf",
	"outputs.tf",
}

type Factory interface {
	New(ctx context.Context, cfg *parser.Boundry) (Terraform, error)
	Upload(ctx context.Context, dir string) (Terraform, error)
}

type factory struct {
	*Options
}

func (m *factory) New(ctx context.Context, cfg *parser.Boundry) (Terraform, error) {
	return NewTerraform(ctx, m.Options, cfg)
}

func (m *factory) Upload(ctx context.Context, dir string) (Terraform, error) {
	return nil, nil
}

func NewFactory(ctx context.Context, cfg *config.Config) (Factory, error) {
	opts := []Option{
		WithBaseDir(cfg.BaseDir),
		WithExecPath(cfg.ExecPath),
		WithDryRun(cfg.DryRun),
	}

	if len(cfg.AwsS3Bucket) > 0 && len(cfg.AwsS3Region) > 0 {
		opts = append(opts, WithBackendS3(cfg.AwsS3Bucket, cfg.AwsS3Region))
	}

	o := NewOptions()
	for _, opt := range opts {
		opt(o)
	}

	return &factory{
		Options: o,
	}, nil
}
