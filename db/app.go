package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

//データベースへの接続
func dbConnect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("DBに接続できませんでした")
	}
	return db
}

//データベースの初期化
func dbInit() {
	db := dbConnect()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}
