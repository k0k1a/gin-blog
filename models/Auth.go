package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	result := db.Where(Auth{Username: username, Password: password}).First(&auth)
	return result.RowsAffected > 0
}
