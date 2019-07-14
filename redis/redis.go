package redis

import (
	U "github.com/neurons-platform/gotools/utils"
	"gopkg.in/redis.v5"
	"time"
)

// 对redis操作的封装

var RedisClient *redis.Client

type Redis_client redis.Client

var Rc Redis_client

func (client *Redis_client) ZAdd(key string, score string, member string) bool {
	z := redis.Z{U.Str2Float64(score), member}
	err := RedisClient.ZAdd(key, z).Err()
	return U.Throw(err)
}

func (client *Redis_client) LRange(key string, start, stop int64) []string {
	val, err := RedisClient.LRange(key, start, stop).Result()
	U.Throw(err)
	return val
}

func (client *Redis_client) ZRangeByScore(key string, min string, max string) []string {
	opt := redis.ZRangeBy{min, max, 0, 0}
	val, err := RedisClient.ZRangeByScore(key, opt).Result()
	U.Throw(err)
	return val
}

func (client *Redis_client) Expire(key string, t time.Duration) bool {
	err := RedisClient.Expire(key, t).Err()
	return U.Throw(err)
}

func (client *Redis_client) SAdd(key string, member string) bool {
	err := RedisClient.SAdd(key, member).Err()
	return U.Throw(err)
}

func (client *Redis_client) SMembers(key string) []string {
	val, err := RedisClient.SMembers(key).Result()
	U.Throw(err)
	return val
}

func (client *Redis_client) Keys(key string) []string {
	val, err := RedisClient.Keys(key).Result()
	U.Throw(err)
	return val
}

func (client *Redis_client) Del(key string) bool {
	err := RedisClient.Del(key).Err()
	return U.Throw(err)
}

func (client *Redis_client) Get(key string) string {
	val, err := RedisClient.Get(key).Result()
	U.Throw(err)
	return val
}

func (client *Redis_client) Set(key string, value string, time time.Duration) bool {
	err := RedisClient.Set(key, value, time).Err()
	return U.Throw(err)
}
func (client *Redis_client) HMSet(key string, value map[string]string) bool {
	err := RedisClient.HMSet(key, value).Err()
	return U.Throw(err)
}
func (client *Redis_client) HGet(key string, field string) string {
	val, err := RedisClient.HGet(key, field).Result()
	U.Throw(err)
	return val
}
func (client *Redis_client) HGetAll(key string) map[string]string {
	val, err := RedisClient.HGetAll(key).Result()
	U.Throw(err)
	return val
}

//批量get
func (client *Redis_client) PGetAll(keys []string) map[string]string {
	var get map[string]*redis.StringCmd
	get = make(map[string]*redis.StringCmd)
	pip := RedisClient.Pipeline()
	for _, key := range keys {
		g := pip.Get(key)
		get[key] = g
	}
	pip.Exec()
	var r map[string]string
	r = make(map[string]string)

	for k, v := range get {
		value, _ := v.Result()
		r[k] = value
	}
	return r
}
func (client *Redis_client) HExists(key string, field string) bool {
	val, err := RedisClient.HExists(key, field).Result()
	U.Throw(err)
	return val
}

func NewRedisClient(addr string, password string) *redis.Client {
	var client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})
	return client

}

// 初始化redis客户端
func InitRedisClient(addr string, password string) {
	RedisClient = NewRedisClient(addr, password)
}
