package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/TriLuong/go-sample/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
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

	hashPassword, error := hashPassword(user)
	if error != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error %s", error))
	}
	user.Passoword = hashPassword
	token, error := createJwt("admin@gmail.com")
	user.Token = token
	if error != nil {
		return c.String(http.StatusInternalServerError, "Login ERROR!!!")
	}
	return c.JSON(http.StatusOK, user)
}

func createJwt(id string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        "admin_id",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, error := rawToken.SignedString([]byte("go-sample"))
	if error != nil {
		return "", error
	}
	return token, nil
}

func hashPassword(user models.User) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Passoword), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func GetUsers(c echo.Context) error {
	role := c.QueryParam("role")

	return c.String(http.StatusOK, fmt.Sprintf("Get Users with role %s\n", role))
}

func GetUserById(c echo.Context) error {
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

func AddUser(c echo.Context) error {
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
