package redis

import (
	"offer_tiktok/pkg/constants"
	"time"

	"github.com/go-redis/redis/v7"
)

var (
	expireTime  = time.Hour * 1
	rdbFollows  *redis.Client
	rdbFavorite *redis.Client
)

func InitRedis() {
	rdbFollows = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})
	rdbFavorite = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       1,
	})
}
