package repository

import (
	"net/http"

	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagRepositoryImpl(Db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{Db: Db}
}

// Delete implements TagRepository
func (t *TagRepositoryImpl) Delete(tagId uint) *utils.ResponseError {
	var tag models.Tag
	if err := t.Db.Preload("Posts").First(&tag, tagId).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	if err := t.Db.Model(&tag).Association("Posts").Clear(); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	result := t.Db.Delete(&tag)
	if result.Error != nil {
		return &utils.ResponseError{
			Message: result.Error.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

// GetAll implements TagRepository
func (t *TagRepositoryImpl) GetAll(metadata utils.Metadata) ([]models.Tag, *utils.ResponseError) {
	var tags []models.Tag
	query := t.Db.Preload("Posts")

	if metadata.SearchBy != "" {
		query = query.Where("label LIKE ?", "%"+metadata.SearchBy+"%")
	}

	offset := (metadata.PageNo - 1) * metadata.PageSize
	query = query.Offset(offset).Limit(metadata.PageSize)

	if err := query.Find(&tags).Error; err != nil {
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return tags, nil
}

// GetById implements TagRepository
func (t *TagRepositoryImpl) GetById(tagId uint) (models.Tag, *utils.ResponseError) {
	var tag models.Tag
	if err := t.Db.Preload("Posts").First(&tag, tagId).Error; err != nil {
		return models.Tag{}, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return tag, nil
}

// Create implements TagRepository
func (t *TagRepositoryImpl) Create(tag *models.Tag) *utils.ResponseError {
	for i, post := range tag.Posts {
		var existingPost models.Post
		if err := t.Db.Where("title = ?", post.Title).First(&existingPost).Error; err != nil {
			if err := t.Db.Create(&tag.Posts[i]).Error; err != nil {
				return &utils.ResponseError{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				}
			}
		} else {
			tag.Posts[i] = existingPost
		}
	}

	if err := t.Db.Create(tag).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

// Update implements TagRepository
func (t *TagRepositoryImpl) Update(tag *models.Tag, tagId uint) *utils.ResponseError {
	var existingTag models.Tag
	if err := t.Db.First(&existingTag, tagId).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	for i, post := range tag.Posts {
		var existingPost models.Post
		if err := t.Db.Where("title = ?", post.Title).First(&existingPost).Error; err != nil {
			if err := t.Db.Create(&tag.Posts[i]).Error; err != nil {
				return &utils.ResponseError{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				}
			}
		} else {
			tag.Posts[i] = existingPost
		}
	}

	if err := t.Db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&existingTag).Updates(tag).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
