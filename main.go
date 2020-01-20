package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}
