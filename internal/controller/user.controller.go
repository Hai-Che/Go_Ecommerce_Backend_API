package controller

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/service"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, 20001, result)
}
