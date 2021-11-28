package route

import (
	"go_web_app/controllers"
	"go_web_app/logger"
	"go_web_app/middleware"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}
	r := gin.New()
	// 最重要的就是这个日志库
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册
	r.POST("/register", controllers.RegisterHandler)
	// 登录
	r.POST("/login", controllers.LoginHandler)

	//验证jwt机制
	r.GET("/ping", middleware.JWTAuthMiddleWare(), func(context *gin.Context) {
		// 这里post man 模拟的 将token auth-token
		zap.L().Debug("ping", zap.String("ping-username", context.GetString("username")))
		controllers.ResponseSuccess(context, "pong")
	})

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return r
}
