package article

import (
	"github.com/lidongyooo/swag-blog-api/app/models"
	"github.com/lidongyooo/swag-blog-api/pkg/model"
)

type Tag struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null" json:"name"`

	Articles []*Article `gorm:"many2many:article_tags;" json:"articles"`
}

func GetTagsByNames(names []string) ([]Tag, error) {
	var tags []Tag

	if err := model.DB.Where("name IN ?", names).Find(&tags).Error; err != nil {
		return tags, err
	}

	return tags, nil
}

func TagsCreate(tags []Tag) ([]Tag, error) {
	result := model.DB.Create(&tags)

	var err error
	if err = result.Error; err != nil {
		return tags, err
	}

	return tags, nil
}