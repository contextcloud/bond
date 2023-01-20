package controllers

import (
	"bond/pkg/terra"
	"io/ioutil"
	"net/http"
)

type Deploy interface {
	Plan(w http.ResponseWriter, r *http.Request)
	Apply(w http.ResponseWriter, r *http.Request)
}

type deploy struct {
	terraFactory terra.Factory
}

func (d *deploy) Plan(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	filename := "data.json"
	contentType := r.Header.Get("Content-Type")
	if contentType == "application/hcl" {
		filename = "main.hcl"
	}
	cfg, err := terra.Parse(filename, data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	tf, err := d.terraFactory.New(ctx, cfg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := tf.Plan(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !result {
		// nothing to plan
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// todo: return plan
}

func (d *deploy) Apply(w http.ResponseWriter, r *http.Request) {

}

func NewDeploy(terraFactory terra.Factory) Deploy {
	return &deploy{
		terraFactory: terraFactory,
	}
}
