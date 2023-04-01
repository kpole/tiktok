package redis

import (
	"fmt"
	"strconv"
	"testing"
)

func TestQueryCount(t *testing.T) {
	InitRedis()
	user_id := 1003
	if cnt, err := RdbFollowing.SCard(strconv.Itoa(user_id)).Result(); cnt > 0 {
		// 更新过期时间。
		RdbFollowing.Expire(strconv.Itoa(int(user_id)), ExpireTime)
		fmt.Println(cnt, err)
	}
}

func TestAddFollow(t *testing.T) {
	InitRedis()
	user_id := 1003
	RdbFollowing.SAdd(strconv.Itoa(user_id), 1005)
	RdbFollowing.SAdd(strconv.Itoa(user_id), 1006)
}
