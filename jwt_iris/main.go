package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/kataras/iris"

	_ "github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	// jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

var superAdmin = User{
	1,
	"username",
	"password",
	UserProfile{Name: "myName", Permissions: []string{"super_Admin_Is_All"}},
}

//////////////////////

// User thong tin user
type User struct {
	ID       int
	Username string
	Password string
	Profile  UserProfile
}

// UserProfile thong tin user public
type UserProfile struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// UserClaims tap hop cac yeu cau JWT chua Profile
type UserClaims struct {
	Profile UserProfile `json:"profile"`
	jwt.StandardClaims
}

var signingKey = []byte("signing-key")

var user []User

func signingKeyFn(*jwt.Token) (interface{}, error) {
	return signingKey, nil
}

func main() {
	app := iris.New()

	app.Post("/login", login)

	app.Get("/auth", auth)

	app.Post("/create/{id}", create)

	app.Run(iris.Addr(":8080"))

}

func login(ctx iris.Context) {
	var user User
	err := ctx.ReadJSON(&user)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	username := user.Username
	password := user.Password

	temp := 0

	for _, x := range user {
		if username == x.Username && password == x.Password {
			temp = 1
		}
	}

	if username == superAdmin.Username && password == superAdmin.Password {
		temp = 1
	}

	if temp == 1 {
		claims := UserClaims{
			superAdmin.Profile,
			jwt.StandardClaims{
				Issuer: "test-project",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(signingKey)
		if err != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.WriteString(err.Error())
			ctx.Write([]byte(err.Error()))
			log.Printf("err: %+v\n", err)
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.Write([]byte(ss))
		log.Printf("issued token: %v\n", ss)
		return
	}

	return
}

func auth(ctx iris.Context) {
	x := ctx.Request()
	var claims UserClaims
	token, err := jwtreq.ParseFromRequestWithClaims(x, jwtreq.AuthorizationHeaderExtractor, &claims, signingKeyFn)

	if err != nil {
		log.Println("Failed to parse token")
		return
	}

	if !token.Valid {
		log.Println("Invalid token")
		return
	}
	ctx.StatusCode(iris.StatusOK)
	claimsString := fmt.Sprintf("claims: %v", claims)
	ctx.Write([]byte(claimsString))
	log.Println(claimsString)
}

func create(ctx iris.Context) {
	var newuser User
	ID := ctx.Params().Get("id")

	err := ctx.ReadJSON(&newuser)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	newuser.ID, _ = strconv.Atoi(ID)
	x := ctx.Request()
	y := ctx.ResponseWriter()
	_ = json.NewDecoder(x.Body).Decode(&user)
	user = append(user, newuser)
	json.NewEncoder(y).Encode(newuser)
}
