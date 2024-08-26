package global

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/logger"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)
