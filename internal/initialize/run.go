package initialize

import (
	"fmt"

	"github.com/Hai-Che/Go_Ecommerce_Backend_API/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	m := global.Config.MySQL
	fmt.Println("Test loading config:", m.Host, m.Dbname)
	InitLogger()
	global.Logger.Info("Config log ok!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run("localhost:8002")
}
