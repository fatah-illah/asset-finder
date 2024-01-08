package models

import "gorm.io/gorm"

type PostTag struct {
	TagID  uint `gorm:"primaryKey"`
	PostID uint `gorm:"primaryKey"`

	Post Post `gorm:"foreignKey:PostID"`
	Tag  Tag  `gorm:"foreignKey:TagID"`
}

func (PostTag) TableName() string {
	return "post_tags"
}

func AutoMigratePostTag(db *gorm.DB) {
	err := db.AutoMigrate(&PostTag{})
	if err != nil {
		return
	}
}
