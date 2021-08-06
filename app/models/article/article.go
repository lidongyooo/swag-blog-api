package article

import (
	"github.com/lidongyooo/swag-blog-api/app/models"
	"github.com/lidongyooo/swag-blog-api/pkg/config"
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

func GetArticlesByTag(tagId uint64, page int) (int64, []Article, error) {
	var (
		articles []Article
		tag Tag
		count int64
		pageLimit = config.Viper.GetInt("PAGINATION_LIMIT")
		pageOffset = (page - 1) * pageLimit
	)

	if tagId > 0 {
		model.DB.First(&tag, tagId)
		count = model.DB.Model(&tag).Association("Articles").Count()
		if err := model.DB.Model(&tag).Limit(pageLimit).Offset(pageOffset).Association("Articles").Find(&articles); err != nil {
			return count, articles, err
		}

		var articleIds []uint64
		for _, article := range articles {
			articleIds = append(articleIds, article.Id)
		}
		if err := model.DB.Preload("Tags").Where("id IN ?", articleIds).Order("id DESC").Find(&articles).Error; err != nil {
			return count, articles, err
		}
	} else {
		model.DB.Model(&articles).Count(&count)
		if err := model.DB.Preload("Tags").Limit(pageLimit).Offset(pageOffset).Order("id DESC").Find(&articles).Error; err != nil {
			return count, articles, err
		}
	}

	return count, articles, nil
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