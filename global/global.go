package global

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/logger"
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
