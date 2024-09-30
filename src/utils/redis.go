package utils

import (
	"blogGo/conf"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RC *redis.Client

func InitRedis(ctx context.Context) {
	RC = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.CFG.RedisConf.Host, conf.CFG.RedisConf.Port),
		Password: conf.CFG.RedisConf.Password,
		DB:       conf.CFG.RedisConf.DB,
	})

	if pong, err := RC.Ping(ctx).Result(); err != nil {
		Logging.Fatalf("Fail to connect redis: %v", err)
	} else {
		Logging.Info("Redis connected:", pong)
	}
}
