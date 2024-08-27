package initialize

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/global"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// r.use() logger
	// r.use() cross
	// r.use() limiter
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User
	mainGroup := r.Group("/v1/2024")
	{
		mainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		managerRouter.InitUserRouter(mainGroup)
		managerRouter.InitAdminRouter(mainGroup)
	}

	return r
}
