package article

import (
	"github.com/lidongyooo/swag-blog-api/app/models"
)

type Tag struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null"`

	Articles []*Article `gorm:"many2many:article_tags;"`
}