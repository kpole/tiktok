package redis

import "testing"

var rdb Follows

func TestFollows_AddFollow(t *testing.T) {
	InitRedis()
	rdb.AddFollower(1001, 1002)
}

func TestFollows_AddFollower(t *testing.T) {
	InitRedis()
	rdb.AddFollow(1002, 1001)
}
