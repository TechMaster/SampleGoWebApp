package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Genre struct {
	// tableName is an optional field that specifies custom table name and alias.
	// By default go-pg generates table name and alias from struct name.
	tableName struct{} `sql:"genres,alias:genre"` // default values are the same

	Id     int // Id is automatically detected as primary key
	Name   string
	Rating int `sql:"-"` // - is used to ignore field

	Books []Book `pg:"many2many:book_genres"` // many to many relation

	ParentId  int
	Subgenres []Genre `pg:"fk:parent_id"` // fk specifies prefix foreign key
}

type Image struct {
	Id   int
	Path string
}

type Author struct {
	ID    int     // both "Id" and "ID" are detected as primary key
	Name  string  `sql:",unique"`
	Books []*Book // has many relation

	AvatarId int
	Avatar   Image
}

func (a Author) String() string {
	return fmt.Sprintf("Author<ID=%d Name=%q>", a.ID, a.Name)
}

type BookGenre struct {
	tableName struct{} `sql:"alias:bg"` // custom table alias

	BookId  int `sql:",pk"` // pk tag is used to mark field as primary key
	Book    *Book
	GenreId int `sql:",pk"`
	Genre   *Genre

	Genre_Rating int // belongs to and is copied to Genre model
}

type Book struct {
	Id        int
	Title     string
	AuthorID  int
	Author    Author // has one relation
	EditorID  int
	Editor    *Author   // has one relation
	CreatedAt time.Time `sql:"default:now()"`
	UpdatedAt time.Time

	Genres       []Genre       `pg:"many2many:book_genres"` // many to many relation
	Translations []Translation // has many relation
	Comments     []Comment     `pg:"polymorphic:trackable_"` // has many polymorphic relation
}

func (b *Book) BeforeInsert(db orm.DB) error {
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now()
	}
	return nil
}

// BookWithCommentCount is like Book model, but has additional CommentCount
// field that is used to select data into it. The use of `pg:",override"` tag
// is essential here and it overrides internal model properties such as table name.
type BookWithCommentCount struct {
	Book `pg:",override"`

	CommentCount int
}

type Translation struct {
	tableName struct{} `sql:",alias:tr"` // custom table alias

	Id     int
	BookId int    `sql:"unique:book_id_lang"`
	Book   *Book  // has one relation
	Lang   string `sql:"unique:book_id_lang"`

	Comments []Comment `pg:",polymorphic:trackable_"` // has many polymorphic relation
}

type Comment struct {
	TrackableId   int    // Book.Id or Translation.Id
	TrackableType string // "Book" or "Translation"
	Text          string
}

func initializeDatabase(db *pg.DB) {
	var genre Genre
	var image Image
	var author Author
	var book Book
	var bookgenre BookGenre
	var bookWithCommentCount BookWithCommentCount
	var translation Translation
	var comment Comment

	// Tạo bảng
	for _, model := range []interface{}{&genre, &image, &author, &book, &bookgenre, &bookWithCommentCount, &translation, &comment} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          false,
			FKConstraints: true,
			IfNotExists:   true,
		})
		if err != nil {
			panic(err)
		}
	}

}

func ConnectToDb(user, password, database, addr string) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
		Addr:     addr,
	})
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	if n == 1 {
		fmt.Println("Connected")
	}
	return db
}

func main() {

	db := ConnectToDb("postgres", "123456", "book", "localhost:5432")
	defer db.Close()
	initializeDatabase(db)

	// Insert example

	example := Image {
		Id:   123,
		Path: "pathexample",
	}
	err := db.Insert(&example)
	if err != nil {
		panic(err)
	}
}
