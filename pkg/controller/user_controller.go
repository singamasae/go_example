package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"my_go/pkg/models"
	"my_go/pkg/service"
	"net/http"
)

//TODO mapping API Endpoint to Service Layer
type userController struct {
	userService service.UserService
}

type UserController interface {
	Save(c *gin.Context)
	FindUserById(c *gin.Context)
	FindAll(c *gin.Context)
}

func InitUserController(s service.UserService) UserController {
	return userController{userService: s}
}

func (u userController) Save(c *gin.Context) {
	//TODO implement me
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	user, err := u.userService.Save(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": "Error has occurred"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

func (u userController) FindUserById(c *gin.Context) {
	//TODO implement me
	log.Print("param id: " + c.Param("id"))
	user := u.userService.FindById(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"success": true, "data": user})
}

func (u userController) FindAll(c *gin.Context) {
	//TODO implement me
	users := u.userService.FindAll()
	c.JSON(http.StatusOK, gin.H{"success": true, "data": users})
}
