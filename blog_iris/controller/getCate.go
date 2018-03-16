package controller

import (
	"fmt"
	"../models"
	"github.com/kataras/iris"
) 

func GetCate(ctx iris.Context) {
	data := model.Getcategory(db)
	ctx.View("category.html", data)
	fmt.Println(data)
}