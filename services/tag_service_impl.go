package services

import (
	"errors"
	"github.com/fatah-illah/asset-finder/data/request"
	"github.com/fatah-illah/asset-finder/data/response"
	"github.com/fatah-illah/asset-finder/models"
	"github.com/fatah-illah/asset-finder/repositories"
	"github.com/fatah-illah/asset-finder/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

type TagServiceImpl struct {
	TagRepository repositories.TagRepository
	Validator     *validator.Validate
}

func NewTagServiceImpl(TagRepository repositories.TagRepository, validator *validator.Validate) TagServices {
	return &TagServiceImpl{
		TagRepository: TagRepository,
		Validator:     validator,
	}
}

func (s *TagServiceImpl) CreateTag(createRequest request.TagRequest) error {
	if err := s.Validator.Struct(createRequest); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	tag := &models.Tag{Label: createRequest.Label}

	if err := s.TagRepository.Create(tag); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *TagServiceImpl) GetAllTags(metadata utils.Metadata) ([]response.TagResponse, error) {
	tags, err := s.TagRepository.GetAll(metadata)
	if err != nil {
		return nil, err
	}

	var tagsResponse []response.TagResponse
	for _, tag := range tags {
		tagResponse := response.TagResponse{
			ID:    tag.ID,
			Label: tag.Label,
			Posts: mapPostsToResponse(tag.Posts),
		}
		tagsResponse = append(tagsResponse, tagResponse)
	}

	return tagsResponse, nil
}

func (s *TagServiceImpl) GetTagByID(tagID uint) (*response.TagResponse, error) {
	tag, err := s.TagRepository.GetByID(tagID)
	if err != nil {
		return nil, err
	}

	tagResponse := response.TagResponse{
		ID:    tag.ID,
		Label: tag.Label,
		Posts: mapPostsToResponse(tag.Posts),
	}

	return &tagResponse, nil
}

func (s *TagServiceImpl) UpdateTagByID(tagID uint, updateRequest *request.TagRequest) error {
	if err := s.Validator.Struct(updateRequest); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	existingTag, err := s.TagRepository.GetByID(tagID)
	if err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		}
	}

	existingTag.Label = updateRequest.Label

	tagUpdateRequest := &request.TagRequest{
		Label: existingTag.Label,
	}

	if err := s.TagRepository.Update(tagID, tagUpdateRequest); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *TagServiceImpl) DeleteTagByID(tagID uint) error {
	if err := s.TagRepository.Delete(tagID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &utils.ResponseError{
				Message: err.Error(),
				Status:  http.StatusNotFound,
			}
		}
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func mapPostsToResponse(posts []*models.Post) []response.PostResponse {
	var postsResponse []response.PostResponse
	for _, post := range posts {
		postResponse := response.PostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
		}
		postsResponse = append(postsResponse, postResponse)
	}
	return postsResponse
}
