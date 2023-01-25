package tests

import (
	"bond/controllers"
	"bond/pkg/terra"
	"embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/afero"
)

//go:embed data
var content embed.FS

func TestDeploy(t *testing.T) {
	t.Run("can run plan", func(t *testing.T) {
		f, err := content.Open("data/s3/s3.def")
		if err != nil {
			t.Fatalf("failed to open plan.json: %v", err)
		}
		req := httptest.NewRequest("POST", "/deploy/plan", f)
		res := httptest.NewRecorder()

		fs := afero.NewOsFs()
		factory, err := terra.NewFactory(fs, "./output")
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
