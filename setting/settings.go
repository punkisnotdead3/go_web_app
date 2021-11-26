package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Config = new(AppConfig)

// Init 加载配置文件
func Init(filePath string) error {
	// 如果设置filepath就直接使用 否则用当前目录的
	if len(filePath) == 0 {
		viper.SetConfigName("config") // 配置文件的名称
		viper.SetConfigType("yaml")   // 配置文件的扩展名，这里除了json还可以有yaml等格式
		// 这个配置可以有多个，主要是告诉viper 去哪个地方找配置文件
		// 我们这里就是简单配置下 在当前工作目录下 找配置即可
		viper.AddConfigPath(".")
	} else {
		viper.SetConfigFile(filePath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper init failed:", err)
		return err
	}
	// 变化就在这里 有个序列化对象的过程
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Println("viper Unmarshal err", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
	})
	return err
}

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	StartTime    string `mapstructure:"start_time"`
	MachineId    int64  `mapstructure:"machineID"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host              string `mapstructure:"host"`
	Port              int    `mapstructure:"port"`
	User              string `mapstructure:"user"`
	Password          string `mapstructure:"password"`
	DbName            string `mapstructure:"dbname"`
	MaxOpenConnection int    `mapstructure:"max_open_connection"`
	MaxIdleConnection int    `mapstructure:"max_idle_connection"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"passowrd"`
	Post     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
