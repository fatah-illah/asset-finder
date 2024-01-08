package models

type Tag struct {
	ID    uint   `gorm:"primaryKey"`
	Label string `json:"label" gorm:"unique"`
	Posts []Post `gorm:"many2many:post_tags;"`
}
