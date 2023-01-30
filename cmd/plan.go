package cmd

import (
	"bond/handler"
	"bond/pkg/parser"
	"context"
	"fmt"
	"os"

	"github.com/contextcloud/graceful/config"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Plan the configuration",
	Long:  `plan plans the configuration.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		cfg, err := config.NewConfig(ctx)
		if err != nil {
			return err
		}

		terraFactory, err := handler.NewTerraFactory(ctx, cfg)
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

			result, err := tf.Plan(ctx)
			if err != nil {
				return err
			}

			fmt.Printf("Has changes: %v", result)
		}

		cancel()
		<-ctx.Done()
		return nil
	},
}
