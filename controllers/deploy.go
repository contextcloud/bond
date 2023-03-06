package controllers

import (
	"bond/pkg/client"
	"bond/pkg/parser"
	"io/ioutil"
	"net/http"
)

type Deploy interface {
	Apply(w http.ResponseWriter, r *http.Request)
}

type deploy struct {
	p             parser.Parser
	clientFactory client.Factory
}

func (d *deploy) Apply(w http.ResponseWriter, r *http.Request) {
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
	cfg, err := d.p.Parse(filename, data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	client, err := d.clientFactory.New(ctx, cfg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := client.Apply(ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// todo: return output.
}

func NewDeploy(clientFactory client.Factory) Deploy {
	p := parser.NewParser()

	return &deploy{
		p:             p,
		clientFactory: clientFactory,
	}
}
