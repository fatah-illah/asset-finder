package controllers

import (
	services2 "github.com/fatah-illah/asset-finder/services"
)

type ManagerControllers struct {
	PostController
	TagController
}

func NewControllersManager(serviceMgr *services2.ManagerServices) *ManagerControllers {
	return &ManagerControllers{
		*NewPostController(serviceMgr.PostService),
		*NewTagController(serviceMgr.TagServices),
	}
}
