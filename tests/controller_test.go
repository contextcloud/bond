package tests

import (
	"bond/config"
	"bond/controllers"
	"bond/examples"
	"bond/pkg/client"
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeploy(t *testing.T) {
	testdata := []struct {
		Name string
	}{{
		Name: "buckets.hcl",
	}}

	ctx := context.Background()
	cfg := &config.Config{
		BaseDir: "./tmp",
	}
	factory, err := client.NewFactory(ctx, cfg)
	if err != nil {
		t.Fatalf("failed to create factory: %v", err)
		return
	}

	for _, d := range testdata {
		t.Run(d.Name, func(t *testing.T) {
			raw, err := examples.ReadFile(d.Name)
			if err != nil {
				t.Fatalf("failed to open: %v", err)
				return
			}
			reader := bytes.NewReader(raw)

			req := httptest.NewRequest("POST", "/deploy/apply", reader)
			res := httptest.NewRecorder()

			c := controllers.NewDeploy(factory)
			c.Apply(res, req)

			// Check the status code is what we expect.
			if status := res.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			// Check the response body is what we expect.
			expected := `{"alive": true}`
			if res.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					res.Body.String(), expected)
			}
		})
	}
}
