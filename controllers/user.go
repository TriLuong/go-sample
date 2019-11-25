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
