package client

import "context"

type Client interface {
	Apply(ctx context.Context) error
	Destroy(ctx context.Context) error
}
