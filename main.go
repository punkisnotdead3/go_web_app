package main

import (
	"fmt"
	"go_web_app/dao/mysql"
	"go_web_app/dao/redis"
	"go_web_app/logger"
	"go_web_app/route"
	"go_web_app/setting"

	"go.uber.org/zap"
)

func main() {
	// 加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("init settings failed:%s \n", err)
		return
	}
	// 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("init settings failed:%s \n", err)
		return
	}
	zap.L().Debug("logger init success")
	defer zap.L().Sync()
	// 不要遗漏2个 db的close
	// 初始化mysql
	if err := mysql.Init(); err != nil {
		fmt.Printf("init mysql failed:%s \n", err)
		return
	}
	zap.L().Debug("mysql init success")

	// 初始化redis
	if err := redis.Init(); err != nil {
		fmt.Printf("init redis failed:%s \n", err)
		return
	}
	zap.L().Debug("redis init success")
	// 注册路由
	route.Setup()
	// 启动服务 （优雅关机）

}
