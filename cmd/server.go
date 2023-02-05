package cmd

import (
	"bond/config"
	"bond/handler"
	"context"

	"github.com/contextcloud/graceful"
	gconfig "github.com/contextcloud/graceful/config"
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

		cfg, err := config.NewConfig()
		if err != nil {
			return err
		}

		terraFactory, err := config.NewTerraFactory(ctx, cfg)
		if err != nil {
			return err
		}

		handler, err := handler.NewHandler(ctx, terraFactory)
		if err != nil {
			return err
		}

		gcfg, err := gconfig.NewConfig(ctx)
		if err != nil {
			return err
		}

		startable, err := srv.NewStartable(gcfg.SrvAddr, handler)
		if err != nil {
			return err
		}

		tracer, err := srv.NewTracer(ctx, gcfg)
		if err != nil {
			return err
		}

		multi := srv.NewMulti(
			tracer,
			srv.NewMetricsServer(gcfg.MetricsAddr),
			srv.NewHealth(gcfg.HealthAddr),
			startable,
		)

		// graceful?
		graceful.Run(ctx, multi)
		cancel()

		<-ctx.Done()
		return nil
	},
}
