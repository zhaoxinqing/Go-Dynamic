package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client              // redis
var rdbClients map[int]*redis.Client // 声明
var redisDBs = []int{0, 1, 2}

// InitRedis ...
func InitRedis() {

	// rdbClients 初始化
	rdbClients = make(map[int]*redis.Client)
	for _, db := range redisDBs {
		rdb := redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       db,
		})
		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			panic(fmt.Sprintf("Redis %d 库连接失败!!", db))
		}
		rdbClients[db] = rdb

	}
	Redis = rdbClients[redisDBs[0]] // 配置初始化默认 redis db

	fmt.Println("Redis 已连接!!!")
}

// ObtainTargetRedisClient ...
func ObtainTargetRedisClient(db int) *redis.Client {

	// 根据 redcs db 获取目标 redis client
	if db < len(redisDBs) {
		if rdbClient, ok := rdbClients[redisDBs[db]]; ok {
			return rdbClient
		}
	}

	// 默认 redis client db
	return rdbClients[redisDBs[0]]
}
