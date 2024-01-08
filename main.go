package main

import (
	"os"

	"github.com/fatah-illah/asset-finder/config"
	"github.com/fatah-illah/asset-finder/server"
	"github.com/fatah-illah/asset-finder/utils"
	"github.com/rs/zerolog/log"
)

// @title Asset Finder Services API
// @version 1.0
// @description A Post, Tag, and Many-to-Many Relationship between Post and Tag services API in Go using Gin framework.
// @host localhost:8000
// @BasePath /api
func main() {
	utils.SetupTimezone()

	log.Info().Msg("Starting Asset Finder - API Development")

	log.Info().Msg("Initializing configuration ...")
	confHandler := config.InitConfig(getConfigFileName())

	log.Info().Msg("Initializing database ...")
	dbHandler := server.InitDatabase(confHandler)
	log.Info().Msgf("Database initialized. Handler: %+v", dbHandler)

	log.Info().Msg("Initializing HTTP Server ...")
	httpServer := server.InitHttpServer(confHandler, dbHandler)

	httpServer.Start()
}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "asset_finder" + env
	}

	return "asset_finder"
}
