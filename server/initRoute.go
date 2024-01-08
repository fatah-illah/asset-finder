package server

import (
	"net/http"

	"github.com/fatah-illah/asset-finder/controllers"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoute(mgrController *controllers.ManagerControllers) *gin.Engine {
	r := gin.Default()

	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Welcome Home!")

		log.Info().Msg("Request to home endpoint")
	})

	// Setup Swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRouter := r.Group("/api")
	postRouter := baseRouter.Group("/posts")
	tagsRouter := baseRouter.Group("/tags")
	postTagsRouter := baseRouter.Group("/postTags")

	// router (API) end-point Post
	postRouter.GET("", mgrController.GetPosts)
	postRouter.GET("/:postId", mgrController.GetPost)
	postRouter.POST("", mgrController.CreatePost)
	postRouter.PUT("/:postId", mgrController.UpdatePost)
	postRouter.DELETE("/:postId", mgrController.DeletePost)

	// router (API) end-point Tag
	tagsRouter.GET("", mgrController.GetTags)
	tagsRouter.GET("/:tagId", mgrController.GetTag)
	tagsRouter.POST("", mgrController.CreateTag)
	tagsRouter.PUT("/:tagId", mgrController.UpdateTag)
	tagsRouter.DELETE("/:tagId", mgrController.DeleteTag)

	// router (API) end-point PostTags
	postTagsRouter.GET("", mgrController.GetPostTags)
	postTagsRouter.GET("/post/:postId", mgrController.GetPostTagsByPostID)
	postTagsRouter.GET("/tag/:tagId", mgrController.GetPostTagsByTagID)
	postTagsRouter.DELETE("/post/:postId", mgrController.DeletePostTagsByPostID)
	postTagsRouter.DELETE("/tag/:tagId", mgrController.DeletePostTagsByTagID)

	return r
}
