package main

import (
	"encoding/json"
	"fmt"
	"log"
	_"strconv"
	"reflect"

	"github.com/kataras/iris"
	uuid "github.com/satori/go.uuid"

	_ "github.com/dgrijalva/jwt-go" 
	jwt "github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	// jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

var superAdmin = User{
	uuid.Must(uuid.NewV4()),
	"admin",
	"password",
	UserProfile{Name: "AdminName", Permissions: []string{"FullPermissions"}},
}

var user01 = User{
	uuid.Must(uuid.NewV4()),
	"user01",
	"password",
	UserProfile{Name: "UserName01", Permissions: []string{"SomePermissions"}},
}

//////////////////////
 
// User thong tin user
type User struct {
	ID       uuid.UUID
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

	user = append (user, superAdmin)
	user = append (user, user01)

	app.Post("/user/login", login)

	app.Get("/user/verify", verify)

	app.Post("/user/register", register)

	app.Get("/user", alluser)

	app.Delete("/user/{id}", delete)

	app.Run(iris.Addr(":8080"))

}

func login(ctx iris.Context) {
	var userlogin User
	err := ctx.ReadJSON(&userlogin)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	username := userlogin.Username
	password := userlogin.Password

	isUserAuthenticated := 0
	var profile UserProfile
	// duyet qua danh sach cac user hien co
	for _, x := range user {
		if username == x.Username && password == x.Password {
			isUserAuthenticated = 1
			profile = x.Profile
		}
	}

	if isUserAuthenticated == 1 {
		claims := UserClaims{
			profile,
			jwt.StandardClaims{
				Issuer: "test-project",
				//ExpiresAt:
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err := token.SignedString(signingKey)
		if err != nil {	
			ctx.WriteString(err.Error())
			ctx.Write([]byte(err.Error()))
			log.Printf("err: %v\n", err)
			return
		}
		ctx.StatusCode(iris.StatusOK)
		ctx.Write([]byte(tokenStr))
		log.Printf("issued token: %v\n", tokenStr)
		return
	}

	return
}

func verify(ctx iris.Context) {
	x := ctx.Request()
	var claims UserClaims
	token, err := jwtreq.ParseFromRequestWithClaims(x, jwtreq.AuthorizationHeaderExtractor, &claims, signingKeyFn)

	if err != nil {
		log.Println("Failed to parse token")
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())	
		return
	}

	if !token.Valid {
		log.Println("Invalid token")
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString(err.Error())
		return
	}
	claimsString := fmt.Sprintf("%v", claims)
	ctx.Write([]byte(claimsString))
	log.Println(claimsString)
}

func register(ctx iris.Context) {
	var newuser User
	err := ctx.ReadJSON(&newuser)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	u1 := uuid.Must(uuid.NewV4())
	newuser.ID = u1
	x := ctx.Request()
	y := ctx.ResponseWriter()
	_ = json.NewDecoder(x.Body).Decode(&newuser)
	user = append(user, newuser)
	json.NewEncoder(y).Encode(newuser)
}

func alluser(ctx iris.Context) {
	x := ctx.Request()
	var claims UserClaims
	token, err := jwtreq.ParseFromRequestWithClaims(x, jwtreq.AuthorizationHeaderExtractor, &claims, signingKeyFn)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		log.Println("Failed to parse token")
		return
	}

	if !token.Valid {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString(err.Error())
		log.Println("Invalid token")
		return
	}

	fullpermis := []string {"FullPermissions"}
	permis := claims.Profile.Permissions
	if reflect.DeepEqual(fullpermis, permis) {
		fmt.Println("OK")
		ctx.StatusCode(iris.StatusOK)
		claimsString := fmt.Sprintf("%v", claims)
		ctx.JSON(user)
		log.Println(claimsString)
	} else {
		fmt.Println("None")
		ctx.StatusCode(iris.StatusBadRequest)
	}
}

func delete(ctx iris.Context) {
	isUser := ctx.Params().Get("id")
	i, _ := uuid.FromString(isUser)
	for index, item := range user {
		if item.ID == i {
			user = append(user[:index], user[index+1:]...)
			ctx.JSON(user)
			break
		} else {
			ctx.StatusCode(iris.StatusNotFound)
			log.Println("ID Not Found")
		}
	}
	
}