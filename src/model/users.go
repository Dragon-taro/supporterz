package model

import "github.com/jinzhu/gorm"

type User struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func LoadUsers(db *gorm.DB) (*[]User, error) {
	u := new([]User)
	// SELECT * FROM users;
	q := db.Find(&u)
	if err := q.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func LoadUser(db *gorm.DB, id int) (*User, error) {
	u := new(User)
	// SELECT * from users WHERE id = ?;
	q := db.Where("id = ?", id).First(&u)
	if err := q.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func AddUser(db *gorm.DB, name, email string) error {
	u := User{Name: name, Email: email}
	// INSERT INTO users(name, email) values(name, email)
	q := db.Create(&u)
	if err := q.Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, id int, name, email string) error {
	u := new(User)
	// UPDATE FROM users SET name = name, email = email where id = ?
	q := db.Model(&u).Where("id = ?", id).Updates(User{Name: name, Email: email})
	if err := q.Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, id int) error {
	// DELETE FROM users WHERE id = ?
	q := db.Where("id = ?", id).Delete(User{})
	if err := q.Error; err != nil {
		return err
	}
	return nil
}
