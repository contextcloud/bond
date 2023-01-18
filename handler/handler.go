package handler

import (
	"bond/controllers"
	"context"
	"net/http"
	"time"

	"github.com/contextcloud/graceful/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
}

// NewHandler creates a new http handler
func NewHandler(ctx context.Context, c *config.Config) (http.Handler, error) {
	cfg := &Config{}
	if err := c.Parse(cfg); err != nil {
		return nil, err
	}

	// setup the deploy controller
	d := controllers.NewDeploy(nil)

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
