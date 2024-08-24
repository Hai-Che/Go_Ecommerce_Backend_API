package routers

import (
	c "github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/controller"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware())
	v1 := r.Group("/v1")
	{
		v1.GET("/user/1", c.NewUserController().GetUserByID)
	}
	return r
}
