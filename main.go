package main

import (
	"fmt"

	"github.com/TriLuong/go-sample/controllers"
	"github.com/TriLuong/go-sample/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Welcome Server!!!")
	database.MongoConnect()
	e := echo.New()

	e.POST("/auth/login", controllers.Login)

	g := e.Group("/users")
	g.Use(middleware.JWT([]byte("go-sample")))
	g.GET("", controllers.GetUsers)
	g.POST("", controllers.AddUser)
	g.GET("/:id", controllers.GetUserById)

	fmt.Println("Start server")
	e.Start(":5000")
}
