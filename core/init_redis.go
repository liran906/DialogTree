// Path: ./core/init_redis.go

package core

import (
	"dialogTree/global"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis(quiet bool) *redis.Client {
	r := global.Config.Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     r.Addr,     // 不写默认就是这个
		Password: r.Password, // 密码
		DB:       r.DB,       // 默认是0
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		logrus.Fatalln("redis connection error: ", err)
	}
	if !quiet {
		logrus.Infof("Redis [%s] connection successful", global.Config.Redis.Addr)
	}
	return redisDB
}
