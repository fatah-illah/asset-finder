package models

type Post struct {
	ID      uint   `gorm:"primary_key;auto_increment" json:"id"`
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Tags    []*Tag `gorm:"many2many:post_tags" json:"tags"`
}
