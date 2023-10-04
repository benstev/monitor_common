package initializers

import (
	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog/log"
)

// InitializeEnvs intializes envy
func InitializeEnvs() {
	if err := envy.Load("./.env"); err != nil {
		log.Warn().Err(err).Msg("cannot load .env file")
		envy.Reload()
	}
}
