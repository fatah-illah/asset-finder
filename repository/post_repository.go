package repository

import (
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
)

type PostRepository interface {
	Create(post *models.Post) *utils.ResponseError
	Update(post *models.Post, postId uint) *utils.ResponseError
	Delete(postId uint) *utils.ResponseError
	GetById(postId uint) (models.Post, *utils.ResponseError)
	GetAll(metadata utils.Metadata) ([]models.Post, *utils.ResponseError)
}
