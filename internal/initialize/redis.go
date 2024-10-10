package initialize

import (
	"context"
	"fmt"

	"github.com/hoangnguyen/demo-sqlc/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	red := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%v", red.Host, red.Port),
		Password: red.Password,
		DB: red.Database,
		PoolSize: red.PoolSize,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error: ", zap.Error(err))
	}

	global.Logger.Info("Initializing Redis successfully")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting varible: ", zap.Error(err))
		return
	}
	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis getting value: ", zap.Error(err))
		return
	}
	global.Logger.Info("Value of score is: ", zap.String("score", value))
}