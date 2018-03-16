package model

import (
	"fmt"

	"github.com/go-pg/pg"
)

func GetNumberPostsbyCate (db *pg.DB) (data []Categorie) {
	_ = db.Model(&data).Select()
	fmt.Println(data)
	return data
}
func Insertcategory (db *pg.DB, name string) (data Categorie) {
	data = Categorie{Name: name}
	err := db.Insert(&data)
	if err != nil {
		panic(err)
	}
	return data
}

func Updatecategory(db *pg.DB, count int32, name string) (data Categorie) {
	_, _ = db.Model(&data).
	Set("count = ?", count).
	Where("name = ?", name).Update()
	return data
}

func GetCount(db *pg.DB, name string) (data Categorie, number int32) {
	_ = db.Model(&data).Column("count").Where("name LIKE ?", name).Select()
	fmt.Println(data)
	return data, data.Count
}
func GetcolumnName(db *pg.DB) (data []Categorie) {
	_ = db.Model(&data).Column("name").Select()
	return data
}