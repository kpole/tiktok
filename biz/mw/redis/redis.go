package redis

import (
	"time"
	"github.com/go-redis/redis/v7"
)

var (
	ExpireTime = time.Hour * 24
	RdbFollowing, RdbFollower *redis.Client
)

func InitRedis() {
	// 后续可能需要接入config
	RdbFollowing = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	RdbFollower = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 1,
	})
}