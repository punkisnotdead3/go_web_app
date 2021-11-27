package redis

import (
	"fmt"
	"go_web_app/setting"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化连接
func Init(config *setting.RedisConfig) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: config.Password, // no password set
		DB:       config.Db,
		PoolSize: config.PoolSize, // use default DB

	})

	_, err := rdb.Ping().Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
