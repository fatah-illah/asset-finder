package services

import (
	"github.com/fatah-illah/asset-finder/data/request"
	"github.com/fatah-illah/asset-finder/data/response"
	"github.com/fatah-illah/asset-finder/utils"
)

type PostService interface {
	CreatePost(createRequest request.PostRequest) error
	GetAllPosts(metadata utils.Metadata) ([]response.PostResponse, error)
	GetPostByID(postID uint) (*response.PostResponse, error)
	UpdatePostByID(postID uint, updateRequest *request.PostRequest) error
	DeletePostByID(postID uint) error
}
