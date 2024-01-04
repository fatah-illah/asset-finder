package repositories

import (
	"github.com/fatah-illah/asset-finder/data/request"
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
)

type TagRepository interface {
	Create(tag *models.Tag) error
	GetAll(metadata utils.Metadata) ([]*models.Tag, error)
	GetByID(tagID uint) (*models.Tag, error)
	Update(tagID uint, updateRequest *request.TagRequest) error
	Delete(tagID uint) error
}
