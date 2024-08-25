package initialize

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/global"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
