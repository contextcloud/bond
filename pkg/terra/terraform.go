package terra

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
)

type Terraform interface {
	Plan(ctx context.Context) (bool, error)
	Apply(ctx context.Context) error
	Output(ctx context.Context) (map[string]tfexec.OutputMeta, error)
	Destroy(ctx context.Context) error
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
	result, err := t.tf.Plan(ctx, tfexec.Out("plan.tfplan"))
	if err != nil {
		return false, err
	}

	if !result {
		return false, nil
	}

	// read it?
	showPlan, err := t.tf.ShowPlanFileRaw(ctx, "plan.tfplan")
	if err != nil {
		return false, err
	}

	fmt.Println(showPlan)
	return true, nil
}

func (t *terraform) Apply(ctx context.Context) error {
	return t.tf.Apply(ctx)
}

func (t *terraform) Output(ctx context.Context) (map[string]tfexec.OutputMeta, error) {
	return t.tf.Output(ctx)
}

func (t *terraform) Destroy(ctx context.Context) error {
	return t.tf.Destroy(ctx)
}

// Show reads the default state path and outputs the state.
// To read a state or plan file, ShowState or ShowPlan must be used instead.
func (t *terraform) Show(ctx context.Context) (*tfjson.State, error) {
	return t.tf.Show(ctx)
}
