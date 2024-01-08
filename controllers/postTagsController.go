package controllers

import (
	"net/http"
	"strconv"

	"github.com/fatah-illah/asset-finder/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostTagController struct {
	DB *gorm.DB
}

func NewPostTagsController(db *gorm.DB) *PostTagController {
	return &PostTagController{DB: db}
}

// GetPostTags godoc
// @Summary Get all post tags
// @Description Get all post tags
// @Tags postTags
// @Accept json
// @Produce json
// @Success 200 {array} models.PostTag
// @Router /postTags [get]
func (h *PostTagController) GetPostTags(c *gin.Context) {
	var postTags []models.PostTag
	if err := h.DB.Find(&postTags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postTags)
}

// GetPostTagsByPostID godoc
// @Summary Get post tags by post ID
// @Description Get post tags by post ID
// @Tags postTags
// @Accept json
// @Produce json
// @Param postId path int true "Post ID"
// @Success 200 {array} models.PostTag
// @Failure 400 {string} string "Invalid PostID"
// @Router /postTags/byPost/{postId} [get]
func (h *PostTagController) GetPostTagsByPostID(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PostID"})
		return
	}

	var postTags []models.PostTag
	if err := h.DB.Where("post_id = ?", postID).Find(&postTags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postTags)
}

// GetPostTagsByTagID godoc
// @Summary Get post tags by tag ID
// @Description Get post tags by tag ID
// @Tags postTags
// @Accept json
// @Produce json
// @Param tagId path int true "Tag ID"
// @Success 200 {array} models.PostTag
// @Failure 400 {string} string "Invalid TagID"
// @Router /postTags/byTag/{tagId} [get]
func (h *PostTagController) GetPostTagsByTagID(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("tagId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TagID"})
		return
	}

	var postTags []models.PostTag
	if err := h.DB.Where("tag_id = ?", tagID).Find(&postTags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postTags)
}

// DeletePostTagsResponse represents the response format for DeletePostTags
type DeletePostTagsResponse struct {
	Message string `json:"message"`
}

// ...

// DeletePostTagsByPostID godoc
// @Summary Delete post tags by post ID
// @Description Delete post tags by post ID
// @Tags postTags
// @Accept json
// @Produce json
// @Param postId path int true "Post ID"
// @Success 200 {object} DeletePostTagsResponse
// @Failure 400 {string} string "Invalid PostID"
// @Router /postTags/byPost/{postId} [delete]
func (h *PostTagController) DeletePostTagsByPostID(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PostID"})
		return
	}

	if err := h.DB.Where("post_id = ?", postID).Delete(&models.PostTag{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeletePostTagsResponse{Message: "PostTags deleted successfully"})
}

// DeletePostTagsByTagID godoc
// @Summary Delete post tags by tag ID
// @Description Delete post tags by tag ID
// @Tags postTags
// @Accept json
// @Produce json
// @Param tagId path int true "Tag ID"
// @Success 200 {object} DeletePostTagsResponse
// @Failure 400 {string} string "Invalid TagID"
// @Router /postTags/byTag/{tagId} [delete]
func (h *PostTagController) DeletePostTagsByTagID(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("tagId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid TagID"})
		return
	}

	if err := h.DB.Where("tag_id = ?", tagID).Delete(&models.PostTag{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeletePostTagsResponse{Message: "PostTags deleted successfully"})
}
