package handler

import (
	"bond/controllers"
	"bond/pkg/terra"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/contextcloud/graceful/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/afero"
)

type Config struct {
	BaseDir string

	S3 *struct {
		Bucket string `mapstructure:"bucket"`
		Region string `mapstructure:"region"`
	} `mapstructure:"s3"`
}

func NewTerraFactory(ctx context.Context, c *config.Config) (terra.Factory, error) {
	cfg := &Config{
		BaseDir: ".",
	}
	if err := c.Parse(cfg); err != nil {
		return nil, err
	}

	env := map[string]string{}
	envKeys := os.Environ()
	for _, k := range envKeys {
		env[k] = os.Getenv(k)
	}

	backend := &terra.Backend{
		Type: terra.BackendTypeLocal,
	}

	if cfg.S3 != nil {
		backend.Type = terra.BackendTypeS3
		backend.Options = &terra.BackendS3{
			Bucket: cfg.S3.Bucket,
			Region: cfg.S3.Region,
		}
	}

	fs := afero.NewOsFs()
	terraFactory, err := terra.NewFactory(ctx, fs, env, cfg.BaseDir, backend)
	if err != nil {
		return nil, err
	}

	return terraFactory, nil
}

// NewHandler creates a new http handler
func NewHandler(ctx context.Context, c *config.Config) (http.Handler, error) {
	terraFactory, err := NewTerraFactory(ctx, c)
	if err != nil {
		return nil, err
	}

	// setup the deploy controller
	d := controllers.NewDeploy(terraFactory)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/deploy/plan", d.Plan)
	r.Post("/deploy/apply", d.Apply)

	return r, nil
}
