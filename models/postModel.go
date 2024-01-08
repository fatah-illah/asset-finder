package models

type Post struct {
	ID      uint   `gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `gorm:"many2many:post_tags;"`
}
