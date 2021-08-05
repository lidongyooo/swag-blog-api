package article

import "github.com/lidongyooo/swag-blog-api/pkg/model"

type ArticleTag struct {
	ArticleId uint64
	TagId uint64
}

func RemoveArticleTagsByArticleIdAndTagsIds(tagIds []uint64, articleId uint64) (rowsAffected int64, err error) {
	result := model.DB.Where("tag_id NOT IN ?", tagIds).Where("article_id = ?", articleId).Delete(&ArticleTag{})
	err = result.Error
	if err != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}

func GetTagIdsByArticleIdAndTagsIds(tagIds []uint64, articleId uint64) ([]uint64, error) {
	var (
		articleTags []ArticleTag
		existTagIds []uint64
	)

	if err := model.DB.Where("tag_id IN ?", tagIds).Where("article_id = ?", articleId).Find(&articleTags).Error; err != nil {
		return existTagIds, err
	}

	for _, articleTag := range articleTags {
		existTagIds = append(existTagIds, articleTag.TagId)
	}

	return existTagIds, nil
}

func ArticleTagsCreates(tagIds []uint64, articleId uint64) ([]ArticleTag, error)  {
	var articleTags []ArticleTag
	for _, tagId := range tagIds {
		articleTags = append(articleTags, ArticleTag{
			TagId: tagId,
			ArticleId: articleId,
		})
	}

	result := model.DB.Create(&articleTags)

	var err error
	if err = result.Error; err != nil {
		return articleTags, err
	}

	return articleTags, nil
}
