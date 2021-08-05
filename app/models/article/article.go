package article

import (
	"github.com/lidongyooo/swag-blog-api/app/models"
	"github.com/lidongyooo/swag-blog-api/pkg/model"
)

type Article struct {
	models.BaseModel

	Title string 	`gorm:"type:varchar(255);not null" binding:"required" json:"title"`
	Body string		`gorm:"not null" binding:"required" json:"body"`

	Tags []*Tag `gorm:"many2many:article_tags;" json:"tags"`
}

func ArticleCreate(article Article) (Article, error) {
	result := model.DB.Create(&article)

	var err error
	if err = result.Error; err != nil {
		return article, err
	}

	return article, nil
}

func ArticleUpdate(article Article) (Article, error) {
	result := model.DB.Save(&article)

	var err error
	if err = result.Error; err != nil {
		return article, err
	}

	return article, nil
}

func Get(id uint64) (Article, error) {
	var article Article

	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}


func GetArticleById(id uint64) (Article, error) {
	var article Article

	if err := model.DB.Preload("Tags").First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}