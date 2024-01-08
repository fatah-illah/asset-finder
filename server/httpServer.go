package server

import (
	"github.com/fatah-illah/asset-finder/controllers"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type HttpServer struct {
	config             *viper.Viper
	router             *gin.Engine
	ManagerControllers controllers.ManagerControllers
}

func InitHttpServer(config *viper.Viper, dbInstance *gorm.DB) HttpServer {
	managerControllers := controllers.NewManagerControllers(dbInstance)

	router := InitRoute(managerControllers)

	return HttpServer{
		config:             config,
		router:             router,
		ManagerControllers: *managerControllers,
	}
}

// Start HttpServer
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))

	if err != nil {
		log.Fatal().Err(err).Msg("Error while starting HTTP Server: %v")
	}
}
