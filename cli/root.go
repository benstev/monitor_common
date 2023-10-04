package cli

import (
	"github.com/benstev/monitor_common/app"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd(builder app.AppBuilder) {
	c := cobra.Command{}

	c.AddCommand(NewServeCmd(builder))

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
