package repository

import (
	"net/http"

	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
	"gorm.io/gorm"
)

type PostTagRepositoryImpl struct {
	Db *gorm.DB
}

func NewPostTagRepositoryImpl(Db *gorm.DB) PostTagRepository {
	return &PostTagRepositoryImpl{Db: Db}
}

// DeleteByTagId implements PostTagRepository
func (pt *PostTagRepositoryImpl) DeleteByTagId(tagId uint) *utils.ResponseError {
	if err := pt.Db.Where("tag_id = ?", tagId).Delete(&models.PostTag{}).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

// DeleteByPostId implements PostTagRepository
func (pt *PostTagRepositoryImpl) DeleteByPostId(postId uint) *utils.ResponseError {
	if err := pt.Db.Where("post_id = ?", postId).Delete(&models.PostTag{}).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return nil
}

// GetByTagId implements PostTagRepository
func (pt *PostTagRepositoryImpl) GetByTagId(tagId uint) ([]models.PostTag, *utils.ResponseError) {
	var postTags []models.PostTag
	if err := pt.Db.Where("tag_id = ?", tagId).Find(&postTags).Error; err != nil {
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return postTags, nil
}

// GetByPostId implements PostTagRepository
func (pt *PostTagRepositoryImpl) GetByPostId(postId uint) ([]models.PostTag, *utils.ResponseError) {
	var postTags []models.PostTag
	if err := pt.Db.Where("post_id = ?", postId).Find(&postTags).Error; err != nil {
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return postTags, nil
}

// GetAll implements PostTagRepository
func (pt *PostTagRepositoryImpl) GetAll(metadata utils.Metadata) ([]models.PostTag, *utils.ResponseError) {
	var postTags []models.PostTag
	query := pt.Db.Model(&models.PostTag{}).Preload("Post").Preload("Tag")

	if metadata.SearchBy != "" {
		searchTerm := "%" + metadata.SearchBy + "%"
		query = query.Joins("JOIN tags ON tags.id = post_tags.tag_id").
			Joins("JOIN posts ON posts.id = post_tags.post_id").
			Where("tags.label LIKE ? OR posts.title LIKE ?", searchTerm, searchTerm)
	}

	offset := (metadata.PageNo - 1) * metadata.PageSize
	query = query.Offset(offset).Limit(metadata.PageSize)

	if err := query.Find(&postTags).Error; err != nil {
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return postTags, nil
}
