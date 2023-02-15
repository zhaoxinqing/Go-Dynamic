package lib

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Redis ...
type RedisConfigs struct {
	Addr     string `yaml:"addr" mapstructure:"addr"`
	Password string `yaml:"password" mapstructure:"password"`
	DBs      []int  `yaml:"dbs" mapstructure:"dbs"`
	DBQueue  int    `yaml:"db_queue"  mapstructure:"db_queue"`
}

var rdbClients map[int]*redis.Client // 声明
var defaultRedisDB = 0

// InitRedis ...
func InitRedis(conf *RedisConfigs) (*redis.Client, error) {
	var (
		ctx      = context.Background()
		redisDBs = conf.DBs
	)
	// rdbClients 初始化
	// rdbClients := make(map[int]*redis.Client)
	// for _, db := range redisDBs {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       redisDBs[0],
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	// rdbClients[db] = rdb

	// 配置初始化默认 redis db
	// if defaultRedisDB == db {
	// 	Redis = rdb
	// }
	// }
	return rdb, nil
}

// ObtainTargetRedisClient ...
func ObtainTargetRedisClient(db int) *redis.Client {
	// 根据 redcs db 获取目标 redis client
	if db < len(rdbClients) {
		if rdbClient, ok := rdbClients[db]; ok {
			return rdbClient
		}
	}
	// 默认 redis client db
	return rdbClients[defaultRedisDB]
}
