package repository

import (
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
)

type PostTagRepository interface {
	DeleteByTagId(tagId uint) *utils.ResponseError
	DeleteByPostId(postId uint) *utils.ResponseError
	GetByTagId(tagId uint) ([]models.PostTag, *utils.ResponseError)
	GetByPostId(postId uint) ([]models.PostTag, *utils.ResponseError)
	GetAll(metadata utils.Metadata) ([]models.PostTag, *utils.ResponseError)
}
