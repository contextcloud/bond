package controllers

import (
	"bond/pkg/terra"
	"net/http"
)

type Deploy interface {
	Plan(w http.ResponseWriter, r *http.Request)
	Apply(w http.ResponseWriter, r *http.Request)
}

type deploy struct {
	terraManager terra.Manager
}

func (d *deploy) Plan(w http.ResponseWriter, r *http.Request) {
	// parse the file.
}

func (d *deploy) Apply(w http.ResponseWriter, r *http.Request) {

}

func NewDeploy(terraManager terra.Manager) Deploy {
	return &deploy{
		terraManager: terraManager,
	}
}
