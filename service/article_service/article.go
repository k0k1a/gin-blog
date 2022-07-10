package article_service

import (
	"encoding/json"
	"github.com/k0k1a/go-gin-example/models"
	"github.com/k0k1a/go-gin-example/pkg/gredis"
	"github.com/k0k1a/go-gin-example/pkg/logging"
	"github.com/k0k1a/go-gin-example/pkg/setting"
	"github.com/k0k1a/go-gin-example/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

// Get 获取article，先从缓存中获取，缓存中没有再查询数据库
func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exist(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, article, setting.RedisSetting.ExpireTime)
	return article, nil
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleById(a.ID)
}

func (a *Article) Count() (int64, error) {
	return models.GetArticleTotal(a.GetMaps())
}

func (a *Article) GetAll() ([]models.Article, error) {
	var articles, cacheArticles []models.Article
	cache := cache_service.Article{
		TagID: a.TagID,
		State: a.State,

		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	if gredis.Exist(key) {
		data, err := gredis.Get(key)
		if err == nil {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
		logging.Info(err)
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.GetMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, articles, setting.RedisSetting.ExpireTime)
	return articles, nil
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"create_by":       a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	if err := models.AddArticle(article); err != nil {
		return err
	}
	return nil
}

func (a *Article) ExistById() (bool, error) {
	return models.ExistArticleById(a.ID)
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID,map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

func(a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

func (a *Article) GetMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.State != -1 {
		maps["state"] = a.State
	}

	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}
	return maps
}
