package controller

import (
	"fmt"
	"strconv"

	"../models"
	"github.com/kataras/iris"
)
// func GetPublishPost(ctx iris.Context) {
// 	data := model.Getpublished(db)
// 	for i ,x:= range data {
// 		if len(x.Categories) > 0 {
// 			data[i].Categories = data[i].Categories[:1]
// 		}
// 	}
// 	ctx.View("preview.html", data)
// } 
func DeletePost(ctx iris.Context) {
	temp := ctx.Params().Get("id")
	i, _ := strconv.Atoi(temp)
	id := int16(i)
	data := model.Getbyid(db, id)
	for _, values := range data.Categories {
		_, count := model.Get(db, values)
		count = count - 1
		model.Updatecategory(db, count, values)
	}
	fmt.Println(data.Categories)
	model.Deletedata(db, id)
	ctx.HTML("<a href='/post' style='font-size: 18px; font-family: monospace'>/All Posts</a>")
	ctx.HTML("<h1>Deleted !!!</h1>")
	
}
