package model

import (
	_"fmt"

	"github.com/go-pg/pg"
)

func Insertcategory (db *pg.DB, name string) (data Categorie) {
	data = Categorie{Name: name}
	err := db.Insert(&data)
	if err != nil {
		panic(err)
	}
	return data
}