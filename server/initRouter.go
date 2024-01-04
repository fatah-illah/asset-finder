package server

import (
	"net/http"

	"github.com/fatah-illah/asset-finder/controllers"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(controllerMgr *controllers.ManagerControllers) *gin.Engine {
	router := gin.Default()

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Welcome Home!")

		log.Info().Msg("Request to home endpoint")
	})

	// router swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := router.Group("/api")

	postRouter := baseRouter.Group("/posts")
	tagRouter := baseRouter.Group("/tags")

	// router (API) end-point Post
	postRouter.GET("", controllerMgr.GetAllPosts)
	postRouter.GET("/:postID", controllerMgr.GetPostByID)
	postRouter.POST("", controllerMgr.CreatePost)
	postRouter.PUT("/:postID", controllerMgr.UpdatePostByID)
	postRouter.DELETE("/:postID", controllerMgr.DeletePostByID)

	// router (API) end-point Tag
	tagRouter.GET("", controllerMgr.GetAllTags)
	tagRouter.GET("/:tagID", controllerMgr.GetTagByID)
	tagRouter.POST("", controllerMgr.CreateTag)
	tagRouter.PUT("/:tagID", controllerMgr.UpdateTagByID)
	tagRouter.DELETE("/:tagID", controllerMgr.DeleteTagByID)

	return router
}
