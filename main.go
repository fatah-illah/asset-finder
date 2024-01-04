package main

import (
	"github.com/fatah-illah/asset-finder/config"
	"github.com/fatah-illah/asset-finder/server"
	"github.com/fatah-illah/asset-finder/utils"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	utils.SetupTimezone()

	log.Info().Msg("Starting Asset Finder - API Development")

	log.Info().Msg("Initializing configuration ...")
	confHandler := config.InitConfig(getConfigFileName())

	log.Info().Msg("Initializing database...")
	dbHandler := server.InitDatabase(confHandler)
	log.Info().Msgf("Database initialized. Handler: %+v", dbHandler)

	validate := validator.New()
	log.Info().Msg("Initializing HTTP Server!")
	httpServer := server.InitHttpServer(confHandler, validate, dbHandler)

	httpServer.Start()
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "asset_finder" + env
	}

	return "asset_finder"
}
