package cmd

import (
	"bond/handler"
	"context"

	"github.com/contextcloud/graceful"
	"github.com/contextcloud/graceful/config"
	"github.com/contextcloud/graceful/srv"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  `server starts the server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		cfg, err := config.NewConfig(ctx)
		if err != nil {
			return err
		}

		handler, err := handler.NewHandler(ctx, cfg)
		if err != nil {
			return err
		}

		startable, err := srv.NewStartable(cfg.SrvAddr, handler)
		if err != nil {
			return err
		}

		tracer, err := srv.NewTracer(ctx, cfg)
		if err != nil {
			return err
		}

		multi := srv.NewMulti(
			tracer,
			srv.NewMetricsServer(cfg.MetricsAddr),
			srv.NewHealth(cfg.HealthAddr),
			startable,
		)

		// graceful?
		graceful.Run(ctx, multi)
		cancel()

		<-ctx.Done()
		return nil
	},
}
