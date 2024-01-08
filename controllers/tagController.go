package controllers

import (
	"net/http"

	"github.com/fatah-illah/asset-finder/data/response"
	"github.com/fatah-illah/asset-finder/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagController struct {
	DB *gorm.DB
}

func NewTagController(db *gorm.DB) *TagController {
	return &TagController{DB: db}
}

// GetTags 			godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tag
// @Success			200 {object} response.Response{}
// @Router			/tags [get]
func (h *TagController) GetTags(c *gin.Context) {
	var tags []models.Tag
	if err := h.DB.Preload("Posts").Find(&tags).Error; err != nil {
		webResponse := response.NewErrorResponse(http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	c.Header("Content-Type", "application/json")
	webResponse := response.NewSuccessResponse(tags)
	c.JSON(http.StatusOK, webResponse)
}

// GetTag 				godoc
// @Summary				Get Single tag by id.
// @Param				tagId path string true "update tag by id"
// @Description			Return the tag who's tagId value matches id.
// @Produce				application/json
// @Tags				tag
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
func (h *TagController) GetTag(c *gin.Context) {
	tagId := c.Param("tagId")
	var tag models.Tag
	if err := h.DB.Preload("Posts").First(&tag, tagId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, tag)
}

// CreateTag		godoc
// @Summary			Create tag
// @Description		Save tag data in Db.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (h *TagController) CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, post := range tag.Posts {
		var existingPost models.Post
		if err := h.DB.Where("title = ?", post.Title).First(&existingPost).Error; err != nil {
			h.DB.Create(&tag.Posts[i])
		} else {
			tag.Posts[i] = existingPost
		}
	}

	if err := h.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

// UpdateTag godoc
// @Summary Update a tag by ID
// @Description Update a tag by its ID with associated posts
// @Tags tags
// @Accept json
// @Produce json
// @Param tagId path int true "Tag ID"
// @Param input body models.Tag true "Tag object to update"
// @Success 200 {object} models.Tag
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Tag not found"
// @Router /tags/{tagId} [put]
func (h *TagController) UpdateTag(c *gin.Context) {
	tagId := c.Param("tagId")
	var tag models.Tag
	if err := h.DB.First(&tag, tagId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, post := range tag.Posts {
		var existingPost models.Post
		if err := h.DB.Where("title = ?", post.Title).First(&existingPost).Error; err != nil {
			h.DB.Create(&tag.Posts[i])
		} else {
			tag.Posts[i] = existingPost
		}
	}

	if err := h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tag)
}

// DeleteTagResponse represents the response format for DeleteTag
type DeleteTagResponse struct {
	Status string `json:"status"`
}

// ...

// DeleteTag godoc
// @Summary Delete a tag by ID
// @Description Delete a tag by its ID
// @Tags tags
// @Accept json
// @Produce json
// @Param tagId path int true "Tag ID"
// @Success 200 {object} DeleteTagResponse
// @Failure 404 {string} string "Tag not found"
// @Router /tags/{tagId} [delete]
func (h *TagController) DeleteTag(c *gin.Context) {
	tagId := c.Param("tagId")

	var tag models.Tag
	if err := h.DB.Preload("Posts").First(&tag, tagId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	err := h.DB.Model(&tag).Association("Posts").Clear()
	if err != nil {
		return
	}

	if err := h.DB.Delete(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeleteTagResponse{Status: "success"})
}
