package terra

import (
	"context"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
)

type Terraform interface {
	Plan(ctx context.Context) (bool, error)
}

type terraform struct {
	tf *tfexec.Terraform
}

// Plan executes `terraform plan` with the specified options and waits for it
// to complete.
//
// The returned boolean is false when the plan diff is empty (no changes) and
// true when the plan diff is non-empty (changes present).
//
// The returned error is nil if `terraform plan` has been executed and exits
// with either 0 or 2.
func (t *terraform) Plan(ctx context.Context) (bool, error) {
	return t.tf.Plan(ctx)
}

// Show reads the default state path and outputs the state.
// To read a state or plan file, ShowState or ShowPlan must be used instead.
func (t *terraform) Show(ctx context.Context) (*tfjson.State, error) {
	return t.tf.Show(ctx)
}

func NewTerraform(ctx context.Context, workingDir string) (Terraform, error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(ctx)
	if err != nil {
		return nil, err
	}

	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		return nil, err
	}

	if err := tf.Init(ctx, tfexec.Upgrade(true)); err != nil {
		return nil, err
	}

	return &terraform{
		tf: tf,
	}, nil
}
