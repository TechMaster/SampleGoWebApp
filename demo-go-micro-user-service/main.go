package main

import ( 
	"context"
	"log"

	proto "github.com/SampleGoWebApp/demo-go-micro-user-service/proto"
	"github.com/micro/go-micro"
	_"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)


type Userservice struct{}

func (g *Userservice) Testconnect(ctx context.Context) error {
	log.Printf("Received request")
	return nil
}

func (e *Userservice) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
	/*
		Verify here
	*/

	log.Printf("Received request")
	rsp.Email = "nameservice: " + server.DefaultOptions().Name + ". " + req.Email
	return nil
}

func (e *Userservice) CreateUser(ctx context.Context, req *proto.CreateUserRequest, rsp *proto.StatusResponse) error {
	/*
		Verify here
	*/

	log.Printf("Received request")
	rsp.Code = 404 // example 
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("testconn"),
	)
	service.Init()

	proto.RegisterUserserviceHandler(service.Server(), new(Userservice))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}