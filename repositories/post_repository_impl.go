package repositories

import (
	"errors"
	"github.com/fatah-illah/asset-finder/data/request"
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/services"
	"github.com/fatah-illah/asset-finder/utils"
	"gorm.io/gorm"
	"net/http"
)

type PostRepositoryImpl struct {
	dbHandler *gorm.DB
}

func NewPostRepositoryImpl(dbHandler *gorm.DB) PostRepository {
	return &PostRepositoryImpl{
		dbHandler: dbHandler,
	}
}

func (repo *PostRepositoryImpl) GetDBHandler() *gorm.DB {
	return repo.dbHandler
}

func (repo *PostRepositoryImpl) Create(post *models.Post) error {
	if post == nil {
		return &utils.ResponseError{
			Message: "post is nil",
			Status:  http.StatusBadRequest,
		}
	}

	if len(post.Tags) > 0 {
		var tagsToUpdate []*models.Tag
		tagService := &services.TagService{DbHandler: repo.dbHandler}

		for _, tag := range post.Tags {
			createdTag := tagService.CreateTagIfNotExists(tag.Label)
			tagsToUpdate = append(tagsToUpdate, createdTag)
		}

		post.Tags = tagsToUpdate
	}

	return repo.dbHandler.Create(post).Error
}

func (repo *PostRepositoryImpl) GetAll(metadata utils.Metadata) ([]*models.Post, error) {
	var posts []*models.Post
	offset := (metadata.PageNo - 1) * metadata.PageSize

	query := repo.dbHandler.Offset(offset).Limit(metadata.PageSize)

	if metadata.SearchBy != "" {
		query = query.Where("title LIKE ?", "%"+metadata.SearchBy+"%")
	}

	err := query.Find(&posts).Error
	if err != nil {
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return posts, nil
}

func (repo *PostRepositoryImpl) GetByID(postID uint) (*models.Post, error) {
	var post models.Post
	err := repo.dbHandler.First(&post, postID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &utils.ResponseError{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			}
		}
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &post, nil
}

func (repo *PostRepositoryImpl) Update(postID uint, updateRequest *request.PostRequest) error {
	existingPost, err := repo.GetByID(postID)
	if err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	existingPost.Title = updateRequest.Title
	existingPost.Content = updateRequest.Content

	if len(updateRequest.Tags) > 0 {
		var tagsToUpdate []*models.Tag
		tagService := &services.TagService{DbHandler: repo.dbHandler}

		for _, tagLabel := range updateRequest.Tags {
			tag := tagService.CreateTagIfNotExists(tagLabel)
			tagsToUpdate = append(tagsToUpdate, tag)
		}

		if err := repo.dbHandler.Model(existingPost).Association("Tags").Replace(tagsToUpdate); err != nil {
			return &utils.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	if err := repo.dbHandler.Save(existingPost).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (repo *PostRepositoryImpl) Delete(postID uint) error {
	err := repo.dbHandler.Delete(&models.Post{}, postID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}
	return err
}
