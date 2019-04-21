package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConnect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "user:password@/sample_db?parseTime=true&charset=utf8mb4&interpolateParams=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
