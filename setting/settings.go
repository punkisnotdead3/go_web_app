package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Init 加载配置文件
func Init() error {
	viper.SetConfigName("config") // 配置文件的名称
	viper.SetConfigType("yaml")   // 配置文件的扩展名，这里除了json还可以有yaml等格式
	// 这个配置可以有多个，主要是告诉viper 去哪个地方找配置文件
	// 我们这里就是简单配置下 在当前工作目录下 找配置即可
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper init failed:", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
	})
	return err
}
