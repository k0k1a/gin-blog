package models

import "gorm.io/gorm"

type Article struct {
	Model
	TagID int `json:"tag_id,omitempty" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title,omitempty"`
	Desc          string `json:"desc,omitempty"`
	Content       string `json:"content,omitempty"`
	CreatedBy     string `json:"created_by,omitempty"`
	ModifiedBy    string `json:"modified_by,omitempty"`
	State         int    `json:"state,omitempty"`
	CoverImageUrl string `json:"cover_image_url,omitempty"`
}

func ExistArticleById(id int) (bool, error) {
	result := db.First(&Article{}, id)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func GetArticleTotal(maps interface{}) (count int64, err error) {
	err = db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return
}

func GetArticles(pageNum, pageSize int, maps interface{}) (articles []Article, err error) {
	err = db.Preload("Tag").Limit(pageSize).Offset(pageNum).Where(maps).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Preload("Tag").First(&article, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["desc"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) error {
	if err := db.Where("id=?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
