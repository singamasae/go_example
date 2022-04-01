package main

import (
	"github.com/gin-gonic/gin"
	"my_go/pkg/controller"
	"my_go/pkg/db"
	"my_go/pkg/repository"
	"my_go/pkg/service"
)

func main() {
	router := gin.Default()

	db, _ := db.Connect()
	userRepository := repository.InitUserRepository(db)
	userService := service.InitUserInterface(userRepository)
	userController := controller.InitUserController(userService)

	users := router.Group("user")
	users.GET("/validate/:id", userController.ValidateUser)
	users.GET("/all", userController.FindAll)

	router.Run("localhost:8000")
}
