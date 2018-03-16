package controller

import (
	"../models"
	"github.com/kataras/iris"
) 

func GetCate(ctx iris.Context) {
	data := model.GetNumberPostsbyCate(db)
	ctx.View("category.html", data)
}