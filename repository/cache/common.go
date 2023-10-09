package cache

import (
	"gin_mall/conf"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"strconv"
)

// RedisClient reids客户端实例
var RedisClient *redis.Client

func InitCache() {
	Redis()
}

// Redis 初始化redis连接
func Redis() {
	db, _ := strconv.ParseUint(conf.Config.Redis.RedisDbName, 10, 64) //64代表uint64
	client := redis.NewClient(&redis.Options{
		Addr: conf.Config.Redis.RedisHost + ":" + conf.Config.Redis.RedisPort,
		//Password:  "",
		DB: int(db),
	})
	// redis连接测试
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
