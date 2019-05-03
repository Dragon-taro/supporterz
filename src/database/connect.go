package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBConnect() (*gorm.DB, error) {
	/**
	mysqlへの接続
	parseTime: 時間をよしなにパースしてくれる
	charset: 文字コードの指定
	interpolateParams: プレースホルダ置換
	*/
	db, err := gorm.Open("mysql", "user:password@/sample_db?parseTime=true&charset=utf8mb4&interpolateParams=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
