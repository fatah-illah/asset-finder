package repository

import (
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
)

type TagRepository interface {
	Create(tag *models.Tag) *utils.ResponseError
	Update(tag *models.Tag, tagId uint) *utils.ResponseError
	Delete(tagId uint) *utils.ResponseError
	GetById(tagId uint) (models.Tag, *utils.ResponseError)
	GetAll(metadata utils.Metadata) ([]models.Tag, *utils.ResponseError)
}
