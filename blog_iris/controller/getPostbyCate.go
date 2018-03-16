package controller

import (
	"fmt"
	"../models"
	"github.com/kataras/iris"
) 
func GetPostbyCate(ctx iris.Context) {
	url := ctx.Params().Get("category")
	data := model.Getcolumncategories(db)
	// for _, x:= range data {
	// 	fmt.Println (x) 
	// }
	data2 := model.GetbyCate(db ,url)

	fmt.Println ("-------------------------------------------")
	fmt.Println (url)
	fmt.Println ("-------")
	fmt.Println (data)
	fmt.Println ("---")
	fmt.Println (data2)
} 