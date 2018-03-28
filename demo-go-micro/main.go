package main

import (
	"context"
	"fmt"
	"log"

	proto "./proto"
	pg "github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-micro/cmd"
	"github.com/satori/go.uuid"
	_ "golang.org/x/crypto/openpgp/packet"
)

type db struct {
	DB *pg.DB
}

const (
	dbHost     = "localhost:5432"
	dbName     = "postgres"
	dbPassword = "123"
	dbUser     = "postgres"
)

// Userservice struct
type Userservice struct{}

// CreateUser from user.pb.go
func (e *Userservice) CreateUser(ctx context.Context, req *proto.CreateUserRequest, rsp *proto.StatusResponse) error {
	log.Printf("Received request")
	
	var db *pg.DB
	// Ket noi toi database
	db = pg.Connect(&pg.Options{
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
		Addr:     dbHost,
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	if n == 1 {
		fmt.Println("Connected Database")
		Createtable(db)
	}
	// cùng 1 đoạn code, trên main vẫn có thể Insert vào Database
	id := uuid.Must(uuid.NewV4())
	x := proto.User{Id: id.String(), FirstName: "12345"}
	err = db.Insert(&x)
	if err != nil {
		panic(err)
	}

	rsp.Code = 200 // example
	return nil
}

// Hello : test connect
func (e *Userservice) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

//=======================================================================================//
func main() {
	

	// new service ----------------------------------------------
	service := micro.NewService(
		micro.Name("newservice"),
	)
	service.Init()

	proto.RegisterUserserviceHandler(service.Server(), new(Userservice))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// Createtable trong database neu chua co
func Createtable(db *pg.DB) {
	var x proto.User
	for _, model := range []interface{}{&x} {
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
