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

type PostServiceImpl struct {
	PostRepository repositories.PostRepository
	Validator      *validator.Validate
	dbHandler      *gorm.DB
}

func NewPostServiceImpl(postRepository repositories.PostRepository, validator *validator.Validate, dbHandler *gorm.DB) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		Validator:      validator,
		dbHandler:      dbHandler,
	}
}

func (s *PostServiceImpl) CreatePost(createRequest request.PostRequest) error {
	if err := s.Validator.Struct(createRequest); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	var tags []*models.Tag
	tagService := &TagService{DbHandler: s.PostRepository.GetDBHandler()}

	for _, tagLabel := range createRequest.Tags {
		tag := tagService.CreateTagIfNotExists(tagLabel)
		tags = append(tags, tag)
	}

	post := &models.Post{
		Title:   createRequest.Title,
		Content: createRequest.Content,
		Tags:    tags,
	}

	if err := s.PostRepository.Create(post); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (s *PostServiceImpl) GetAllPosts(metadata utils.Metadata) ([]response.PostResponse, error) {
	posts, err := s.PostRepository.GetAll(metadata)
	if err != nil {
		return nil, err
	}

	var postsResponse []response.PostResponse
	for _, post := range posts {
		postResponse := response.PostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			Tags:    mapTagsToResponse(post.Tags),
		}
		postsResponse = append(postsResponse, postResponse)
	}

	return postsResponse, nil
}

func (s *PostServiceImpl) GetPostByID(postID uint) (*response.PostResponse, error) {
	post, err := s.PostRepository.GetByID(postID)
	if err != nil {
		return nil, err
	}

	postResponse := response.PostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		Tags:    mapTagsToResponse(post.Tags),
	}

	return &postResponse, nil
}

func (s *PostServiceImpl) UpdatePostByID(postID uint, updateRequest *request.PostRequest) error {
	if err := s.Validator.Struct(updateRequest); err != nil {
		return &utils.ResponseError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
	}

	existingPost, err := s.PostRepository.GetByID(postID)
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
		tagService := &TagService{DbHandler: s.dbHandler}

		for _, tagLabel := range updateRequest.Tags {
			tag := tagService.CreateTagIfNotExists(tagLabel)
			tagsToUpdate = append(tagsToUpdate, tag)
		}

		existingPost.Tags = tagsToUpdate

		updatedRequest := &request.PostRequest{
			Title:   existingPost.Title,
			Content: existingPost.Content,
			Tags:    make([]string, len(existingPost.Tags)),
		}
		for i, tag := range existingPost.Tags {
			updatedRequest.Tags[i] = tag.Label
		}

		if err := s.PostRepository.Update(postID, updatedRequest); err != nil {
			return &utils.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
		}
	}

	return nil
}

func (s *PostServiceImpl) DeletePostByID(postID uint) error {
	if err := s.PostRepository.Delete(postID); err != nil {
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

func mapTagsToResponse(tags []*models.Tag) []response.TagResponse {
	var tagsResponse []response.TagResponse
	for _, tag := range tags {
		tagResponse := response.TagResponse{
			ID:    tag.ID,
			Label: tag.Label,
		}
		tagsResponse = append(tagsResponse, tagResponse)
	}
	return tagsResponse
}
