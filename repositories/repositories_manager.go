package repositories

import (
	"gorm.io/gorm"
)

type ManagerRepositories struct {
	PostRepository
	TagRepository
}

func NewRepositoriesManager(dbHandler *gorm.DB) *ManagerRepositories {
	return &ManagerRepositories{
		PostRepository: NewPostRepositoryImpl(dbHandler),
		TagRepository:  NewTagRepositoryImpl(dbHandler),
	}
}
