package initialize

import (
	"fmt"

	"github.com/Hai-Che/Go_Ecommerce_Backend_API/global"
)

func Run() {
	LoadConfig()
	m := global.Config.MySQL
	fmt.Println("Test loading config:", m.Host, m.Dbname)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()
	r.Run("localhost:8002")
}
