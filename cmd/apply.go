package cmd

import (
	"bond/config"
	"bond/pkg/parser"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the configuration",
	Long:  `apply applies the configuration.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		cfg, err := config.NewConfig()
		if err != nil {
			return err
		}

		terraFactory, err := config.NewTerraFactory(ctx, cfg)
		if err != nil {
			return err
		}

		p := parser.NewParser()

		for _, filename := range args {
			data, err := os.ReadFile(filename)
			if err != nil {
				return err
			}

			boundry, err := p.Parse(filename, data)
			if err != nil {
				return err
			}
			tf, err := terraFactory.New(ctx, boundry)
			if err != nil {
				return err
			}

			if err := tf.Apply(ctx); err != nil {
				return err
			}

			output, err := tf.Output(ctx)
			if err != nil {
				return err
			}

			raw, err := json.Marshal(output)
			if err != nil {
				return err
			}

			fmt.Printf("%s", string(raw))
		}

		cancel()
		<-ctx.Done()
		return nil
	},
}
