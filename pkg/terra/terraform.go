package terra

import (
	"bond/pkg/parser"
	"bond/pkg/utils"
	"context"
	"fmt"
	"os"

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
	tf     *tfexec.Terraform
	dryRun bool
}

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
	if t.dryRun {
		return nil
	}

	return t.tf.Apply(ctx)
}

func (t *terraform) Output(ctx context.Context) (map[string]tfexec.OutputMeta, error) {
	return t.tf.Output(ctx)
}

func (t *terraform) Destroy(ctx context.Context) error {
	if t.dryRun {
		return nil
	}
	return t.tf.Destroy(ctx)
}

func (t *terraform) Show(ctx context.Context) (*tfjson.State, error) {
	return t.tf.Show(ctx)
}

func NewTerraform(ctx context.Context, opts *Options, cfg *parser.Boundry) (Terraform, error) {
	workingDir, err := ensure(opts.Fs, opts.BaseDir, cfg.Id)
	if err != nil {
		return nil, err
	}

	tf, err := tfexec.NewTerraform(workingDir, opts.ExecPath)
	if err != nil {
		return nil, err
	}

	// TODO toggle this on and off
	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	env := utils.MergeMaps(opts.Env, cfg.Env)
	if err := tf.SetEnv(env); err != nil {
		return nil, err
	}

	// write the provider again this time with remote backend
	if err := createProviders(opts.Fs, workingDir, cfg, opts.BackendType, opts.BackendOptions); err != nil {
		return nil, err
	}
	if err := createMain(opts.Fs, workingDir, cfg.Resources); err != nil {
		return nil, err
	}
	if err := createOutputs(opts.Fs, workingDir, cfg.Resources); err != nil {
		return nil, err
	}

	if err := tf.Init(ctx, tfexec.Upgrade(true)); err != nil {
		return nil, err
	}

	return &terraform{
		tf:     tf,
		dryRun: opts.DryRun,
	}, nil
}
