package cmd

import (
	"bond/config"
	"bond/pkg/client"
	"context"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload files to the server",
	Long:  `upload uploads files to the server.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		cfg, err := config.NewConfig()
		if err != nil {
			return err
		}

		factory, err := client.NewFactory(ctx, cfg)
		if err != nil {
			return err
		}

		for _, arg := range args {
			client, err := factory.Upload(ctx, arg)
			if err != nil {
				return err
			}

			if err := client.Apply(ctx); err != nil {
				return err
			}
		}

		cancel()
		<-ctx.Done()
		return nil
	},
}
