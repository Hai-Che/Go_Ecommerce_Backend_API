package initialize

import (
	"context"
	"fmt"

	"github.com/Hai-Che/Go_Ecommerce_Backend_API/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,
	})
	_, err := rdb.Ping(ctx).Result()
	checkErrorPanic(err, "Failed to init Redis db")

	fmt.Println("Redis connect successfully")
	global.Rdb = rdb
	RedisTesting()
}

func RedisTesting() {
	err := global.Rdb.Set(ctx, "test", 100, 0).Err()
	checkErrorPanic(err, "Failed to set redis variable")
	value, err := global.Rdb.Get(ctx, "test").Result()
	checkErrorPanic(err, "Failed to get redis variable")
	global.Logger.Info(`Value is:`, zap.String("test", value))
}
