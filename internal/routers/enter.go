package routers

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/routers/manager"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
