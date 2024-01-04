package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupTimezone() {
	wib, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal().Err(err).Msg("Error setting timezone")
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02|15:04:05|Z07:00",
	})

	log.Info().Str("timezone", wib.String()).Msg("Timezone set successfully,")
}
