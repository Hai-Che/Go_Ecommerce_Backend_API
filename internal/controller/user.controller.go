package controller

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/service"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/response"
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
	response.SuccessResponse(c, 20001, []string{"1", "2", "3"})
	response.ErrorResponse(c, 20003)
}
