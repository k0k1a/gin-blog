package models

type Article struct {
	Model
	TagID int `json:"tag_id,omitempty" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title,omitempty"`
	Desc       string `json:"desc,omitempty"`
	Content    string `json:"content,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	State      int    `json:"state,omitempty"`
}

func ExistArticleById(id int) bool {
	result := db.First(&Article{}, id)
	return result.RowsAffected > 0
}

func GetArticleTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Limit(pageSize).Offset(pageNum).Where(maps).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	db.Preload("Tag").First(&article, id)
	return
}

func EditArticle(id int, data interface{}) bool {
	result := db.Model(&Article{}).Where("id=?", id).Updates(data)
	return result.RowsAffected > 0
}
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["desc"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	result := db.Where("id=?", id).Delete(&Article{})
	return result.RowsAffected > 0
}
