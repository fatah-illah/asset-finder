package services

import (
	"github.com/fatah-illah/asset-finder/data/request"
	"github.com/fatah-illah/asset-finder/data/response"
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
	"gorm.io/gorm"
)

type TagService struct {
	DbHandler *gorm.DB
}

type TagServices interface {
	CreateTag(createRequest request.TagRequest) error
	GetAllTags(metadata utils.Metadata) ([]response.TagResponse, error)
	GetTagByID(tagID uint) (*response.TagResponse, error)
	UpdateTagByID(tagID uint, updateRequest *request.TagRequest) error
	DeleteTagByID(tagID uint) error
}

func (tagService *TagService) CreateTagIfNotExists(label string) *models.Tag {
	var existingTag models.Tag
	if err := tagService.DbHandler.Where("label = ?", label).First(&existingTag).Error; err != nil {
		newTag := &models.Tag{Label: label}
		tagService.DbHandler.Create(newTag)
		return newTag
	}
	return &existingTag
}
