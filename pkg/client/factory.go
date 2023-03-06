package client

import (
	"bond/config"
	"bond/pkg/parser"
	"bond/pkg/terra"
	"context"
)

type Factory interface {
	New(ctx context.Context, cfg *parser.Boundry) (Client, error)
	Upload(ctx context.Context, dir string) (Client, error)
}

type factory struct {
	cfg *config.Config

	terraFactory terra.Factory
}

func (m *factory) New(ctx context.Context, cfg *parser.Boundry) (Client, error) {
	c, err := m.terraFactory.New(ctx, cfg)
	if err != nil {
		return nil, err
	}
	return &tfclient{
		tf: c,
	}, nil
}

func (m *factory) Upload(ctx context.Context, dir string) (Client, error) {
	return nil, nil
}

func NewFactory(ctx context.Context, cfg *config.Config) (Factory, error) {
	terraFactory, err := terra.NewFactory(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &factory{
		cfg:          cfg,
		terraFactory: terraFactory,
	}, nil
}

type tfclient struct {
	tf terra.Terraform
}

func (tf *tfclient) Apply(ctx context.Context) error {
	changes, err := tf.tf.Plan(ctx)
	if err != nil {
		return err
	}

	if !changes {
		return nil
	}

	return tf.tf.Apply(ctx)
}

func (tf *tfclient) Destroy(ctx context.Context) error {
	return tf.tf.Destroy(ctx)
}
