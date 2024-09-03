package util

import (
	"os"

	log "github.com/rs/zerolog/log"
)

const (
	ENV_IBKR_API_URL       = "IBKR_API_URL"
	ENV_MY_ACCOUNT_INDEX   = "MY_ACCOUNT_INDEX"
	ENV_CORS_ENABLED       = "CORS_ENABLED"
	ENV_MAX_PREEXPIRY_TIME = "MAX_PREEXPIRY_TIME"
)

func Env(envVar string) (val string) {
	val, exists := os.LookupEnv(envVar)
	if !exists {
		log.Fatal().Str("ENV", envVar).Msg("required ENV is not set")
	}
	return
}
