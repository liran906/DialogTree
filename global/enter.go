// Path: ./global/enter.go

package global

import (
	"dialogTree/conf"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
	Redis  *redis.Client
)
