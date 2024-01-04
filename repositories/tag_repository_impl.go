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

type TagRepositoryImpl struct {
	dbHandler *gorm.DB
}

func NewTagRepositoryImpl(dbHandler *gorm.DB) TagRepository {
	return &TagRepositoryImpl{
		dbHandler: dbHandler,
	}
}

func (repo *TagRepositoryImpl) Create(tag *models.Tag) error {
	if tag == nil {
		return &utils.ResponseError{
			Message: "tag is nil",
			Status:  http.StatusBadRequest,
		}
	}

	tagService := &services.TagService{DbHandler: repo.dbHandler}
	createdTag := tagService.CreateTagIfNotExists(tag.Label)

	*tag = *createdTag

	return repo.dbHandler.Create(tag).Error
}

func (repo *TagRepositoryImpl) GetAll(metadata utils.Metadata) ([]*models.Tag, error) {
	var tags []*models.Tag
	offset := (metadata.PageNo - 1) * metadata.PageSize

	query := repo.dbHandler.Offset(offset).Limit(metadata.PageSize)

	if metadata.SearchBy != "" {
		query = query.Where("title LIKE ?", "%"+metadata.SearchBy+"%")
	}

	err := query.Find(&tags).Error
	if err != nil {
		return nil, &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return tags, nil
}

func (repo *TagRepositoryImpl) GetByID(tagID uint) (*models.Tag, error) {
	var tag models.Tag
	err := repo.dbHandler.First(&tag, tagID).Error
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
	return &tag, nil
}

func (repo *TagRepositoryImpl) Update(tagID uint, updateRequest *request.TagRequest) error {
	existingTag, err := repo.GetByID(tagID)
	if err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	existingTag.Label = updateRequest.Label

	if err := repo.dbHandler.Save(existingTag).Error; err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (repo *TagRepositoryImpl) Delete(tagID uint) error {
	err := repo.dbHandler.Delete(&models.Tag{}, tagID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}
	return err
}
