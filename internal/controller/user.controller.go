package controller

import (
	"net/http"

	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInfoUser(),
	})
}
