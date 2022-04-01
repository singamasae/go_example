package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"my_go/pkg/service"
	"net/http"
)

type userController struct {
	userService service.UserServiceInterface
}

type UserControllerInterface interface {
	ValidateUser(*gin.Context)
	FindAll(c *gin.Context)
}

func InitUserController(s service.UserServiceInterface) UserControllerInterface {
	return userController{userService: s}
}

func (u userController) FindAll(c *gin.Context) {
	//TODO implement me
	users := u.userService.FindAll()
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "data": users})
}

func (u userController) ValidateUser(c *gin.Context) {
	//TODO implement me
	log.Print(c.Param("id"))
	var result = u.userService.ValidateUser(c.Param("id"))
	message := ""
	if result == true {
		message = "valid user"
	} else {
		message = "invalid user"
	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "message": message})
}
