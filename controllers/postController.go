package controllers

import (
	"net/http"

	"github.com/fatah-illah/asset-finder/data/response"
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{DB: db}
}

// GetPosts godoc
// @Summary Get a post by ID
// @Description Get a post by its ID with tags
// @Tags posts
// @Accept json
// @Produce json
// @Param postId path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {string} string "Post not found"
// @Router /posts/{postId} [get]
func (h *PostController) GetPosts(c *gin.Context) {
	var posts []models.Post

	if err := h.DB.Preload("Tags").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}

	c.Header("Content-Type", "application/json")

	webResponse := response.NewSuccessResponse(posts)
	c.JSON(http.StatusOK, webResponse)
}

// GetPost godoc
// @Summary Get a post by ID
// @Description Get a post by its ID with tags
// @Tags posts
// @Accept json
// @Produce json
// @Param postId path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {string} string "Post not found"
// @Router /posts/{postId} [get]
func (h *PostController) GetPost(c *gin.Context) {
	postId := c.Param("postId")
	var post models.Post
	if err := h.DB.Preload("Tags").First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with tags
// @Tags posts
// @Accept json
// @Produce json
// @Param input body models.Post true "Post object to create"
// @Success 200 {object} models.Post
// @Failure 400 {string} string "Bad request"
// @Router /posts [post]
func (h *PostController) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, tag := range post.Tags {
		var existingTag models.Tag
		if err := h.DB.Where("label = ?", tag.Label).First(&existingTag).Error; err != nil {
			h.DB.Create(&post.Tags[i])
		} else {
			post.Tags[i] = existingTag
		}
	}

	if err := h.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdatePost godoc
// @Summary Update a post by ID
// @Description Update a post by its ID with tags
// @Tags posts
// @Accept json
// @Produce json
// @Param postId path int true "Post ID"
// @Param input body models.Post true "Post object to update"
// @Success 200 {object} models.Post
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Post not found"
// @Router /posts/{postId} [put]
func (h *PostController) UpdatePost(c *gin.Context) {
	postId := c.Param("postId")
	var post models.Post
	if err := h.DB.First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, tag := range post.Tags {
		var existingTag models.Tag
		if err := h.DB.Where("label = ?", tag.Label).First(&existingTag).Error; err != nil {
			h.DB.Create(&post.Tags[i])
		} else {
			post.Tags[i] = existingTag
		}
	}

	if err := h.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

// DeletePostResponse represents the response format for DeletePost
type DeletePostResponse struct {
	Status string `json:"status"`
}

// ...

// DeletePost godoc
// @Summary Delete a post by ID
// @Description Delete a post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Param postId path int true "Post ID"
// @Success 200 {object} DeletePostResponse
// @Failure 404 {string} string "Post not found"
// @Router /posts/{postId} [delete]
func (h *PostController) DeletePost(c *gin.Context) {
	postId := c.Param("postId")

	var post models.Post
	if err := h.DB.Preload("Tags").First(&post, postId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	err := h.DB.Model(&post).Association("Tags").Clear()
	if err != nil {
		return
	}

	if err := h.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DeletePostResponse{Status: "success"})
}
