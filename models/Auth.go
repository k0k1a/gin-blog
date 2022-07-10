package models

import "gorm.io/gorm"

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	result := db.Where(Auth{Username: username, Password: password}).First(&auth)
	if err := result.Error; err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return result.RowsAffected > 0, nil
}
