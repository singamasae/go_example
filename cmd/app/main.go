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

	dB, _ := db.Connect()
	userRepository := repository.InitUserRepository(dB)
	userService := service.InitUserInterface(userRepository)
	userController := controller.InitUserController(userService)

	users := router.Group("user")
	users.POST("/", userController.Save)
	users.GET("/:id", userController.FindUserById)
	users.GET("/all", userController.FindAll)

	router.Run("localhost:8000")
}
