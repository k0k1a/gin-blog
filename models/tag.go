package models

import "gorm.io/gorm"

type Tag struct {
	Model

	Name       string `json:"name,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	State      int    `json:"state,omitempty"`
}

func GetTags(pageNum int, pageSize int, maps map[string]interface{}) (tags []Tag, err error) {
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error; err != nil {
		return nil, err
	}
	return
}

func GetTagTotal(maps map[string]interface{}) (count int64, err error) {
	if err = db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return
}

//ExistTagByName 根据name判断tag是否存在
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	result := db.Select("id").Where("name=?", name).Find(&tag)
	if err := result.Error; err != nil {
		return false, err
	}
	return result.RowsAffected > 0, nil
}

func ExistTagById(id int) (bool, error) {
	result := db.First(&Tag{}, id)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTag(id int) error {
	if err := db.Delete(&Tag{}, id).Error;err!=nil{
		return err
	}
	return nil
}

// EditTag 修改
func EditTag(id int, data interface{}) error {
	result := db.Model(&Tag{}).Where("id=?", id).Updates(data)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
