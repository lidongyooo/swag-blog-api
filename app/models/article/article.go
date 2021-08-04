package article

import (
	"github.com/lidongyooo/swag-blog-api/app/models"
)

type Article struct {
	models.BaseModel

	Title string 	`gorm:"type:varchar(255);not null" binding:"required" json:"title"`
	Body string		`gorm:"not null" binding:"required" json:"body"`

	Tags []*Tag `gorm:"many2many:article_tags;" json:"tags"`
}

