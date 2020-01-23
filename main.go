package main

import (
	"github.com/shuto/go_todo/db"
	"github.com/shuto/go_todo/route"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	//dbInit()
	db.DbInit()

	//Index
	router.GET("/", route.Index)
	//Create
	router.POST("/new", route.Create)
	//Detail
	router.GET("/detail/:id", route.Detail)
	//Update
	router.POST("/update/:id", route.Update)
	//削除確認
	router.GET("/delete_check/:id", route.DelConf)
	//Delete
	router.POST("/delete/:id", route.Delete)

	router.Run()

}
