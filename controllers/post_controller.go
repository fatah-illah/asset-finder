package controllers

import (
	"net/http"
	"strconv"

	"github.com/fatah-illah/asset-finder/data/request"
	"github.com/fatah-illah/asset-finder/data/response"
	"github.com/fatah-illah/asset-finder/services"
	"github.com/fatah-illah/asset-finder/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{
		postService: service,
	}
}

// GetAllPosts godoc
// @Summary Get all posts
// @Produce application/json
// @Success 200 {object} response.Response{}
// @Router /api/posts [get]
func (controller *PostController) GetAllPosts(ctx *gin.Context) {
	log.Info().Msg("get all posts")

	metadata := utils.Metadata{
		PageNo:   getIntParam(ctx, "pageNo", 1),
		PageSize: getIntParam(ctx, "pageSize", 10),
		SearchBy: getStringParam(ctx, "searchBy", ""),
	}

	postsResponse, err := controller.postService.GetAllPosts(metadata)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	webResponse := response.NewSuccessResponse(postsResponse)
	ctx.JSON(http.StatusOK, webResponse)
}

func getIntParam(ctx *gin.Context, key string, defaultValue int) int {
	value, err := strconv.Atoi(ctx.Query(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func getStringParam(ctx *gin.Context, key, defaultValue string) string {
	value := ctx.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetPostByID godoc
// @Summary Get post by ID
// @Produce application/json
// @Param id path int true "Post ID"
// @Success 200 {object} response.Response{}
// @Router /api/posts/{id} [get]
func (controller *PostController) GetPostByID(ctx *gin.Context) {
	log.Info().Msg("get post by ID")
	postID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	postResponse, err := controller.postService.GetPostByID(uint(postID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	webResponse := response.NewSuccessResponse(postResponse)
	ctx.JSON(http.StatusOK, webResponse)
}

// CreatePost godoc
// @Summary Add a new post
// @Produce application/json
// @Consumes application/json
// @Param post body request.PostRequest true "New Post"
// @Success 200 {object} response.Response{}
// @Router /api/posts [post]
func (controller *PostController) CreatePost(ctx *gin.Context) {
	log.Info().Msg("create a new post")
	createPostRequest := request.PostRequest{}
	err := ctx.ShouldBindJSON(&createPostRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = controller.postService.CreatePost(createPostRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	webResponse := response.NewSuccessResponse(nil)
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdatePostByID godoc
// @Summary Update an existing post by ID
// @Produce application/json
// @Consumes application/json
// @Param id path int true "Post ID"
// @Param post body request.PostRequest true "Updated Post"
// @Success 200 {object} response.Response{}
// @Router /api/posts/{id} [put]
func (controller *PostController) UpdatePostByID(ctx *gin.Context) {
	log.Info().Msg("update an existing post by ID")
	postID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	updatePostRequest := &request.PostRequest{}
	err = ctx.ShouldBindJSON(updatePostRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = controller.postService.UpdatePostByID(uint(postID), updatePostRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	webResponse := response.NewSuccessResponse(nil)
	ctx.JSON(http.StatusOK, webResponse)
}

// DeletePostByID godoc
// @Summary Delete a post by ID
// @Produce application/json
// @Param id path int true "Post ID"
// @Success 200 {object} response.Response{}
// @Router /api/posts/{id} [delete]
func (controller *PostController) DeletePostByID(ctx *gin.Context) {
	log.Info().Msg("delete a post by ID")
	postID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = controller.postService.DeletePostByID(uint(postID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	webResponse := response.NewSuccessResponse(nil)
	ctx.JSON(http.StatusOK, webResponse)
}
