package model

import "github.com/jinzhu/gorm"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func LoadUsers(db *gorm.DB) (*[]User, error) {
	users := new([]User)
	query := db.Find(&users)

	if err := query.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func LoadUser(db *gorm.DB, id int) (*User, error) {
	user := new(User)
	query := db.Where("id = ?", id).First(&user)

	if err := query.Error; err != nil {
		return nil, err
	}

	return user, nil
}
