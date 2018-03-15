package controller

import (
	"../models"
	"github.com/kataras/iris"
) 
func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
func GetPublishPost(ctx iris.Context) {
	data := model.Getpublished(db)
	for i ,x:= range data {
		if len(x.Categories) > 0 {
			data[i].Categories = data[i].Categories[:1]
		}
	}
	ctx.View("preview.html", data)
} 