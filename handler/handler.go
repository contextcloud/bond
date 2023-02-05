package handler

import (
	"bond/controllers"
	"bond/pkg/terra"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewHandler creates a new http handler
func NewHandler(ctx context.Context, terraFactory terra.Factory) (http.Handler, error) {
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
