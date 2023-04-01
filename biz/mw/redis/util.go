package redis

import (
	"strconv"

	"github.com/go-redis/redis/v7"
)

// add k & v to redis
func add(c *redis.Client, k string, v int64) {
	tx := c.TxPipeline()
	tx.SAdd(k, v)
	tx.Expire(k, expireTime)
	tx.Exec()
}

// del k & v
func del(c *redis.Client, k string, v int64) {
	tx := c.TxPipeline()
	tx.SRem(k, v)
	tx.Expire(k, expireTime)
	tx.Exec()
}

// check the set of k if exist
func check(c *redis.Client, k string) bool {
	if e, _ := c.Exists(k).Result(); e > 0 {
		return true
	}
	return false
}

// exist check the relation k and v if exist
func exist(c *redis.Client, k string, v int64) bool {
	if e, _ := c.SIsMember(k, v).Result(); e {
		c.Expire(k, expireTime)
		return true
	}
	return false
}

// count get the size of the set of key
func count(c *redis.Client, k string) (sum int64, err error) {
	if sum, err = c.SCard(k).Result(); err == nil {
		c.Expire(k, expireTime)
		return sum, err
	}
	return sum, err
}

func get(c *redis.Client, k string) (vt []int64) {
	v, _ := c.SMembers(k).Result()
	c.Expire(k, expireTime)
	for _, vs := range v {
		v_i64, _ := strconv.ParseInt(vs, 10, 64)
		vt = append(vt, v_i64)
	}
	return vt
}
