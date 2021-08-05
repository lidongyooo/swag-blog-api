package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/app/models/article"
	res "github.com/lidongyooo/swag-blog-api/pkg/response"
	"github.com/lidongyooo/swag-blog-api/pkg/slices"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type ArticlesController struct {
	BaseController
}

type ArticleExtension struct {
	article.Article

	TagName string `json:"tag_name" binding:"required"`
}

func (articles *ArticlesController) Store (context *gin.Context)  {
	artExt := &ArticleExtension{}
	err := context.ShouldBindJSON(&artExt)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, res.Error(res.INVALID_PARAMS, err))
	} else {
		tagNames := strings.Split(artExt.TagName, ",")
		tagIds, err := tagsHandler(tagNames, 0)
		if err != nil {
			articles.ResponseForError(context, err)
			return
		}

		_article, err := article.ArticleCreate(article.Article{
			Title: artExt.Title,
			Body: artExt.Body,
		})
		if err != nil {
			articles.ResponseForError(context, err)
			return
		}

		_, err = article.ArticleTagsCreates(tagIds, _article.Id)
		if err != nil {
			articles.ResponseForError(context, err)
			return
		}

		_article, err = article.GetArticleById(_article.Id)
		if err != nil {
			articles.ResponseForError(context, err)
			return
		}

		context.JSON(http.StatusCreated, res.New(http.StatusCreated, "ok").WithData(_article))
	}
}

func (articles *ArticlesController) Update (context *gin.Context)  {
	artExt := &ArticleExtension{}
	err := context.ShouldBindJSON(&artExt)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, res.Error(res.INVALID_PARAMS, err))
	} else {
		var _article article.Article
		context.ShouldBindUri(&_article)
		
		_article, err = article.Get(_article.Id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				context.JSON(http.StatusNotFound, res.New(http.StatusNotFound, "未找到此文章"))
			}else {
				articles.ResponseForError(context, err)
			}
			return
		}

		tagNames := strings.Split(artExt.TagName, ",")
		_, err := tagsHandler(tagNames, _article.Id)
		if err != nil {
			articles.ResponseForError(context, err)
			return
		}

		_article.Title = artExt.Title
		_article.Body = artExt.Body
		_article, err = article.ArticleUpdate(_article)
		if err != nil {
			articles.ResponseForError(context, err)
			return
		}

		_article, err = article.GetArticleById(_article.Id)
		if err != nil {
			articles.ResponseForError(context, err)
			return
		}

		context.JSON(http.StatusOK, res.New(http.StatusOK, "ok").WithData(_article))
	}
}


func tagsHandler(tagNames []string, articleId uint64) (tagIds []uint64, err error) {
	tags, err := article.GetTagsByNames(tagNames)
	if err != nil {
		return
	} else {
		for _, tag := range tags {
			if index := slices.SearchSlices(tagNames, tag.Name); index != -1 {
				tagNames = slices.Remove(tagNames, index)
			}
			tagIds = append(tagIds, tag.Id)
		}

		var newTags []article.Tag
		for _, tagName := range tagNames {
			newTags = append(newTags, article.Tag{
				Name: tagName,
			})
		}
		if len(newTags) > 0 {
			newTags, err = article.TagsCreate(newTags)
		}

		if err != nil {
			return
		}

		for _, tag := range newTags{
			tagIds = append(tagIds, tag.Id)
		}

		if articleId > 0 {
			newTagIds, err := article.GetTagIdsByArticleIdAndTagsIds(tagIds, articleId)
			if err != nil {
				return tagIds, err
			}

			_, err = article.RemoveArticleTagsByArticleIdAndTagsIds(newTagIds, articleId)
			if err != nil {
				return tagIds, err
			}

			for _, tagId := range newTagIds {
				if index := slices.SearchSlicesUint64(tagIds, tagId); index != -1 {
					tagIds = slices.RemoveUint64(tagIds, index)
				}
			}
			fmt.Println(tagIds)
			if len(tagIds)  > 0 {
				_, err = article.ArticleTagsCreates(tagIds, articleId)
				if err != nil {
					return tagIds, err
				}
			}
		}

		return
	}
}