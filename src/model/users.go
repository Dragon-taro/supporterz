package model

import "github.com/jinzhu/gorm"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func LoadUsers(db *gorm.DB) (*[]User, error) {
	users := new([]User)
	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
