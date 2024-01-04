package models

import (
	"github.com/fatah-illah/asset-finder/services"
	"gorm.io/gorm"
)

type Tag struct {
	ID    uint    `gorm:"primary_key;auto_increment" json:"id"`
	Label string  `gorm:"type:varchar(255);unique;not null" json:"label"`
	Posts []*Post `gorm:"many2many:post_tags" json:"posts"`
}

func (tag *Tag) CreateTagIfNotExists(db *gorm.DB, label string) *Tag {
	tagService := services.TagService{DbHandler: db}
	return tagService.CreateTagIfNotExists(label)
}
