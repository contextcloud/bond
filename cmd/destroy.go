package cmd

import (
	"context"
	"fmt"
	"os"

	"bond/config"
	"bond/pkg/parser"

	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy resources",
	Long:  `destroy destroys resources.`,
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

			if err := tf.Destroy(ctx); err != nil {
				return err
			}

			fmt.Printf("Destroyed")
		}

		cancel()
		<-ctx.Done()
		return nil
	},
}
