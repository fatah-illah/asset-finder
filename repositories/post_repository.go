package repositories

import (
	"github.com/fatah-illah/asset-finder/data/request"
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	GetAll(metadata utils.Metadata) ([]*models.Post, error)
	GetByID(postID uint) (*models.Post, error)
	Update(postID uint, updateRequest *request.PostRequest) error
	Delete(postID uint) error

	GetDBHandler() *gorm.DB
}
