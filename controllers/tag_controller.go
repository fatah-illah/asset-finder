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

type TagController struct {
	tagService services.TagServices
}

func NewTagController(service services.TagServices) *TagController {
	return &TagController{
		tagService: service,
	}
}

// GetAllTags godoc
// @Summary Get all tags
// @Produce application/json
// @Success 200 {object} response.Response{}
// @Router /api/tags [get]
func (controller *TagController) GetAllTags(ctx *gin.Context) {
	log.Info().Msg("get all tags")

	metadata := utils.Metadata{
		PageNo:   getIntParam(ctx, "pageNo", 1),
		PageSize: getIntParam(ctx, "pageSize", 10),
		SearchBy: getStringParam(ctx, "searchBy", ""),
	}

	tagsResponse, err := controller.tagService.GetAllTags(metadata)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	webResponse := response.NewSuccessResponse(tagsResponse)
	ctx.JSON(http.StatusOK, webResponse)
}

// GetTagByID godoc
// @Summary Get tag by ID
// @Produce application/json
// @Param id path int true "Tag ID"
// @Success 200 {object} response.Response{}
// @Router /api/tags/{id} [get]
func (controller *TagController) GetTagByID(ctx *gin.Context) {
	log.Info().Msg("get tag by ID")
	tagID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	tagResponse, err := controller.tagService.GetTagByID(uint(tagID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	webResponse := response.NewSuccessResponse(tagResponse)
	ctx.JSON(http.StatusOK, webResponse)
}

// CreateTag godoc
// @Summary Add a new tag
// @Produce application/json
// @Consumes application/json
// @Param tag body request.TagRequest true "New Tag"
// @Success 200 {object} response.Response{}
// @Router /api/tags [post]
func (controller *TagController) CreateTag(ctx *gin.Context) {
	log.Info().Msg("create a new tag")
	createTagRequest := request.TagRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = controller.tagService.CreateTag(createTagRequest)
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

// UpdateTagByID godoc
// @Summary Update an existing tag by ID
// @Produce application/json
// @Consumes application/json
// @Param id path int true "Tag ID"
// @Param tag body request.TagRequest true "Updated Tag"
// @Success 200 {object} response.Response{}
// @Router /api/tags/{id} [put]
func (controller *TagController) UpdateTagByID(ctx *gin.Context) {
	log.Info().Msg("update an existing tag by ID")
	tagID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	updateTagRequest := &request.TagRequest{}
	err = ctx.ShouldBindJSON(updateTagRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = controller.tagService.UpdateTagByID(uint(tagID), updateTagRequest)
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

// DeleteTagByID godoc
// @Summary Delete a tag by ID
// @Produce application/json
// @Param id path int true "Tag ID"
// @Success 200 {object} response.Response{}
// @Router /api/tags/{id} [delete]
func (controller *TagController) DeleteTagByID(ctx *gin.Context) {
	log.Info().Msg("delete a tag by ID")
	tagID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}

	err = controller.tagService.DeleteTagByID(uint(tagID))
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
