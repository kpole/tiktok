package redis

import (
	"github.com/go-redis/redis/v7"
	"offer_tiktok/pkg/constants"
	"time"
)

var (
	ExpireTime                = time.Hour * 24
	RdbFollowing, RdbFollower *redis.Client
)

func InitRedis() {
	// 后续可能需要接入config
	RdbFollowing = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
	RdbFollower = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       1,
	})
}
