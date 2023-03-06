package handler

import (
	"bond/controllers"
	"bond/pkg/client"
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewHandler creates a new http handler
func NewHandler(ctx context.Context, clientFactory client.Factory) (http.Handler, error) {
	// setup the deploy controller
	d := controllers.NewDeploy(clientFactory)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/deploy/apply", d.Apply)

	return r, nil
}
