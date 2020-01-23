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

//初期化
func DbInit() {
	db := dbConnect()
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//削除
func DbDelete(id int) {
	db := dbConnect()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//追加
func DbInsert(text string, status string) {
	db := dbConnect()
	db.Create(&Todo{Text: text, Status: status})
}

//更新
func DbUpdate(id int, text string, status string) {
	db := dbConnect()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//全取得
func DbGetAll() []Todo {
	db := dbConnect()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//一つ取得
func DbGetOne(id int) Todo {
	db := dbConnect()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
