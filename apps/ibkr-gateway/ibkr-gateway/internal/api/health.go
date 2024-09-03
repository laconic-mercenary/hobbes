package api

import (
	log "github.com/rs/zerolog/log"
)

func IsHealthy() bool {
	log.Trace().Msg("api.IsHealthy")
	// consider also health checking other resources.. maybe the IBKR API?
	return true
}
