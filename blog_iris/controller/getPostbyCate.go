package controller

import (
	"../models"
	"github.com/kataras/iris"
) 
func GetPostbyCate(ctx iris.Context) {
	url := ctx.Params().Get("category")
	data := model.Getbycategory(db, url)
	for i ,x:= range data {
		if len(x.Categories) > 0 {
			data[i].Categories = data[i].Categories[:1]
		}
	}
	ctx.View("preview.html", data)
}