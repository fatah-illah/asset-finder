package server

import (
	"github.com/fatah-illah/asset-finder/controllers"
	"github.com/fatah-illah/asset-finder/repositories"
	dbContext2 "github.com/fatah-illah/asset-finder/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	ControllersManager controllers.ManagerControllers
}

func InitHttpServer(config *viper.Viper, validator *validator.Validate, dbHandler *gorm.DB) HttpServer {
	repositoriesManager := repositories.NewRepositoriesManager(dbHandler)

	servicesManager := dbContext2.NewServicesManager(repositoriesManager, validator, dbHandler)

	controllersManager := controllers.NewControllersManager(servicesManager)

	router := InitRouter(controllersManager)

	return HttpServer{
		config:             config,
		router:             router,
		ControllersManager: *controllersManager,
	}
}

// Start HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatal().Err(err).Msg("Error while starting HTTP Server: %v")
	}
}
