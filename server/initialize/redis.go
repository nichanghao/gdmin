package initialize

import (
	"context"
	"gitee.com/nichanghao/gdmin/global"
	"github.com/redis/go-redis/v9"
	"log"
)

func InitRedis() {

	redisConfig := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:           redisConfig.Addr,
		Password:       redisConfig.Password,
		DB:             redisConfig.DB,
		MaxActiveConns: redisConfig.MaxActiveConns,
		MinIdleConns:   redisConfig.MinIdleConns,
		MaxIdleConns:   redisConfig.MaxIdleConns,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("redis connect failed, err: %v", err)
	} else {
		log.Println("redis connect success...")
	}

	global.RedisCli = client
}
