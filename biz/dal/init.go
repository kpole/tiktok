package dal

import (
	"offer_tiktok/biz/dal/db"
	"offer_tiktok/biz/mw/redis"
)

// Init init dal
func Init() {
	db.Init() // mysql init
	redis.InitRedis()
}
