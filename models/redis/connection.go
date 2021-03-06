package redis

import (
	"LianFaPhone/lfp-backend-api/tools"
	"fmt"
	"github.com/go-redis/redis"
)

type (
	Redis struct {
		Config *tools.Redis
	}
)

var (
	RedisClient *redis.Client
)

func New(conf *tools.Redis) *Redis {
	return &Redis{
		Config: conf,
	}
}

func (this *Redis) Connection() *redis.Client {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", this.Config.Host, this.Config.Port),
		Password: this.Config.Password,
		DB:       this.Config.Database,
	})

	return RedisClient
}
