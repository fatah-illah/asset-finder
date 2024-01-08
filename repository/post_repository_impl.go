package repository

import (
	"net/http"

	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/utils"
	"gorm.io/gorm"
)

type PostRepositoryImpl struct {
	Db *gorm.DB
}

func NewPostRepositoryImpl(Db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{Db: Db}
}

// Delete implements PostRepository
func (p *PostRepositoryImpl) Delete(postId uint) *utils.ResponseError {
	var post models.Post
	if err := p.Db.Preload("Tags").First(&post, postId).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	if err := p.Db.Model(&post).Association("Tags").Clear(); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	result := p.Db.Delete(&post)
	if result.Error != nil {
		return &utils.ResponseError{
			Message: result.Error.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

// GetAll implements PostRepository
func (p *PostRepositoryImpl) GetAll(metadata utils.Metadata) ([]models.Post, *utils.ResponseError) {
	var posts []models.Post
	query := p.Db.Preload("Tags")

	if metadata.SearchBy != "" {
		query = query.Where("title LIKE ?", "%"+metadata.SearchBy+"%")
	}

	offset := (metadata.PageNo - 1) * metadata.PageSize
	query = query.Offset(offset).Limit(metadata.PageSize)

	if err := query.Find(&posts).Error; err != nil {
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return posts, nil
}

// GetById implements PostRepository
func (p *PostRepositoryImpl) GetById(postId uint) (models.Post, *utils.ResponseError) {
	var post models.Post
	if err := p.Db.Preload("Posts").First(&post, postId).Error; err != nil {
		return models.Post{}, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	return post, nil
}

// Create implements PostRepository
func (p *PostRepositoryImpl) Create(post *models.Post) *utils.ResponseError {
	for i, tag := range post.Tags {
		var existingTag models.Tag
		if err := p.Db.Where("label = ?", post.Title).First(&existingTag).Error; err != nil {
			if err := p.Db.Create(&tag.Posts[i]).Error; err != nil {
				return &utils.ResponseError{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				}
			}
		} else {
			post.Tags[i] = existingTag
		}
	}

	if err := p.Db.Create(post).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

// Update implements PostRepository
func (p *PostRepositoryImpl) Update(post *models.Post, postId uint) *utils.ResponseError {
	var existingPost models.Post
	if err := p.Db.First(&existingPost, postId).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	for i, tag := range post.Tags {
		var existingTag models.Tag
		if err := p.Db.Where("label = ?", tag.Label).First(&existingTag).Error; err != nil {
			if err := p.Db.Create(&tag.Posts[i]).Error; err != nil {
				return &utils.ResponseError{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				}
			}
		} else {
			post.Tags[i] = existingTag
		}
	}

	if err := p.Db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&existingPost).Updates(post).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
