package controller

import (
	_"fmt"
	"time"

	"../models"
	"github.com/go-pg/pg"
	"github.com/kataras/iris"
)

var db *pg.DB
var cate *pg.DB

func CreatePost(ctx iris.Context) {
	if err := ctx.View("create.html"); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
}
func CreatedPost(ctx iris.Context) {
	post := model.Post{}
	err := ctx.ReadForm(&post)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	ctx.ViewData("Title", post.Title)
	ctx.ViewData("Alias", post.Alias)
	ctx.ViewData("Image", post.Image)
	t1 := time.Now().Local().Format(time.RFC822)
	ctx.ViewData("Created_at", t1)
	ctx.ViewData("Published", post.Published)
	ctx.ViewData("Intro_text", post.Intro_text)
	ctx.ViewData("Full_text", post.Full_text)
	ctx.ViewData("Categories", post.Categories) 
	ctx.View("created.html")
	model.Insertdata(db, post.Title, post.Alias, post.Intro_text, post.Full_text, post.Image, post.Published, post.Categories, t1)
	for _, values := range post.Categories {
		_, count := model.Get(db, values)
		count = count + 1
		model.Updatecategory(db, count, values)
	}
}
