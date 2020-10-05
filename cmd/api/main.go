package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/naufalAndika/Majoo-tes/cmd/api/controller"
	"github.com/naufalAndika/Majoo-tes/cmd/api/server"
	"github.com/naufalAndika/Majoo-tes/internal/platform/mysql"
	"github.com/naufalAndika/Majoo-tes/internal/service"
)

func main() {
	e := server.New()

	db, err := mysql.New("root:rahasia123@(localhost:3306)/tes_majoo?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}

	mysql.Seed(db)

	addService(db, e)

	server.Start(e)
}

func addService(db *gorm.DB, e *echo.Echo) {
	userDB := mysql.NewUserDB(db)

	userRouter := e.Group("/users")
	userService := service.NewUser(userDB)

	controller.NewUser(userService, userRouter)

	authRouter := e.Group("/auth")
	authService := service.NewAuth(userDB)

	controller.NewAuth(authService, authRouter)
}
