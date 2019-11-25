package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/TriLuong/go-sample/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	Role      string `json:"role"`
// 	Email     string `json:"email"`
// 	Passoword string `json:"password"`
// 	Phone     string `json:"phone"`
// 	Token     string `json:"token"`
// }

type JwtClaims struct {
	id string `json:"id"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	var user models.User

	defer c.Request().Body.Close()

	b, error := ioutil.ReadAll((c.Request().Body))

	if error != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error %s", error))
	}

	error = json.Unmarshal(b, &user)

	if error != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error %s", error))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Passoword), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	user.Passoword = string(hash)

	token, error := createJwt("admin@gmail.com")
	user.Token = token
	if error != nil {
		return c.String(http.StatusInternalServerError, "Login ERROR!!!")
	}
	return c.JSON(http.StatusOK, user)
}

func createJwt(id string) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			Id:        "admin_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, error := rawToken.SignedString([]byte("go-sample"))
	if error != nil {
		return "", error
	}
	return token, nil
}

func getUsers(c echo.Context) error {
	role := c.QueryParam("role")

	return c.String(http.StatusOK, fmt.Sprintf("Get Users with role %s\n", role))
}

func getUserById(c echo.Context) error {
	id := c.Param("id")

	if id == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Get User with ID %s\n", id))
	}

	if id == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"id": id,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Not found user",
	})

}

func addUser(c echo.Context) error {
	// user := User{}
	var user map[string]interface{}

	defer c.Request().Body.Close()

	b, error := ioutil.ReadAll((c.Request().Body))

	if error != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error %s", error))
	}

	error = json.Unmarshal(b, &user)

	if error != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Cannot convert %s", error))
	}

	return c.JSON(http.StatusOK, user)
}

func main() {
	fmt.Println("Welcome Server!!!")
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.BasicAuth(func(email, password string, c echo.Context) (bool, error) {
	// 	if email == "admin@gmail.com" && password == "123456" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))
	e.POST("/auth/login", login)

	g := e.Group("/users")
	g.Use(middleware.JWT([]byte("go-sample")))
	// g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningMethod: "HS256",
	// 	SigningKey:    []byte("go-sample"),
	// }))

	g.GET("", getUsers)
	g.POST("", addUser)
	g.GET("/:id", getUserById)

	fmt.Println("Start server")
	e.Start(":5000")
}
