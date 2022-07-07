package models

type Tag struct {
	Model

	Name       string `json:"name,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	State      int    `json:"state,omitempty"`
}

func GetTags(pageNum int, pageSize int, maps map[string]interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTotal(maps map[string]interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

//ExistTagByName 根据name判断tag是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	result := db.Select("id").Where("name=?", name).Find(&tag)
	if rows := result.RowsAffected; rows > 0 {
		return true
	}
	return false
}

func ExistTagById(id int) bool {
	result := db.First(&Tag{}, id)
	return result.RowsAffected > 0
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
	return true
}

func DeleteTag(id int) bool {
	db.Delete(&Tag{}, id)
	return true
}

// EditTag 修改
func EditTag(id int, data interface{}) bool {
	result := db.Model(&Tag{}).Where("id=?", id).Updates(data)
	return result.RowsAffected > 0
}
