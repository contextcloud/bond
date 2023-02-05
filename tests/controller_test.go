package tests

import (
	"bond/controllers"
	"bond/examples"
	"bond/pkg/terra"
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
	withBackendS3 := terra.WithBackendS3("contextcloud-bond-test-bucket", "us-east-1")
	withBaseDir := terra.WithBaseDir("./tmp")

	factory, err := terra.NewFactory(ctx, withBackendS3, withBaseDir)
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

			req := httptest.NewRequest("POST", "/deploy/plan", reader)
			res := httptest.NewRecorder()

			c := controllers.NewDeploy(factory)
			c.Plan(res, req)

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
