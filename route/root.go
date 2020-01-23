package route

import (
	"strconv"

	"github.com/shuto/go_todo/db"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	todos := db.DbGetAll()
	ctx.HTML(200, "index.html", gin.H{
		"todos": todos,
	})
}

func Create(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	db.DbInsert(text, status)
	ctx.Redirect(302, "/")
}

func Detail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := db.DbGetOne(id)
	ctx.HTML(200, "detail.html", gin.H{"todo": todo})
}

func Update(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	db.DbUpdate(id, text, status)
	ctx.Redirect(302, "/")
}

func DelConf(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todo := db.DbGetOne(id)
	ctx.HTML(200, "delete.html", gin.H{"todo": todo})
}

func Delete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	db.DbDelete(id)
	ctx.Redirect(302, "/")

}
