package services

import (
	repo "github.com/fatah-illah/asset-finder/repositories"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ManagerServices struct {
	PostService
	TagServices
}

func NewServicesManager(repoMgr *repo.ManagerRepositories, validator *validator.Validate, dbHandler *gorm.DB) *ManagerServices {
	return &ManagerServices{
		PostService: NewPostServiceImpl(repoMgr.PostRepository, validator, dbHandler),
		TagServices: NewTagServiceImpl(repoMgr.TagRepository, validator),
	}
}
