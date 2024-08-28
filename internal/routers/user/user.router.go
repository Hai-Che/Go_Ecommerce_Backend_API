package user

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/controller"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/repo"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (urouter *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/login")
		userRouterPublic.POST("/register", userHandlerNonDependency.Register)
		userRouterPublic.POST("/otp")
	}
	//private
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}

}
