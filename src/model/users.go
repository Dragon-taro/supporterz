package model

import "github.com/jinzhu/gorm"

type User struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func LoadUsers(db *gorm.DB) (*[]User, error) {
	u := new([]User)
	q := db.Find(&u)
	if err := q.Error; err != nil {
		return nil, err
	}

	return u, nil
}

func LoadUser(db *gorm.DB, id int) (*User, error) {
	u := new(User)
	q := db.Where("id = ?", id).First(&u)
	if err := q.Error; err != nil {
		return nil, err
	}

	return u, nil
}

func AddUser(db *gorm.DB, name, email string) error {
	u := User{Name: name, Email: email}
	q := db.Create(&u)
	if err := q.Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, id int, name, email string) error {
	u := new(User)
	q := db.Model(&u).Where("id = ?", id).Updates(User{Name: name, Email: email})
	if err := q.Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, id int) error {
	q := db.Where("id = ?", id).Delete(User{})
	if err := q.Error; err != nil {
		return err
	}
	return nil
}
