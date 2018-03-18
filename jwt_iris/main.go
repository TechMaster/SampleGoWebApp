package main

import (
	"fmt"
	"log"
	_"encoding/json"
	"github.com/kataras/iris"

	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	_"github.com/dgrijalva/jwt-go"
	// jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

// User thong tin user
type User struct {
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

func signingKeyFn(*jwt.Token) (interface{}, error) {
	return signingKey, nil
}

var sampleUser = User{
	"username",
	"password",
	UserProfile{Name: "myName", Permissions: []string{"Admin"}},
}

func main() {
	app := iris.New()

	app.Post("/login", login)

	app.Get("/auth", auth)

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

	if username == sampleUser.Username && password == sampleUser.Password {
		claims := UserClaims{
			sampleUser.Profile,
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

func auth(ctx iris.Context, ) {
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