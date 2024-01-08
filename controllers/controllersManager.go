package controllers

import (
	"gorm.io/gorm"
)

type ManagerControllers struct {
	PostController
	TagController
	PostTagController
}

func NewManagerControllers(dbInstance *gorm.DB) *ManagerControllers {
	return &ManagerControllers{
		*NewPostController(dbInstance),
		*NewTagController(dbInstance),
		*NewPostTagsController(dbInstance),
	}
}
