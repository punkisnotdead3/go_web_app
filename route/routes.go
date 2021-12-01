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

	//v1 版本的路由
	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/register", controllers.RegisterHandler)
	// 登录
	v1.POST("/login", controllers.LoginHandler)

	v1.GET("/community", controllers.CommunityHandler)
	v1.GET("/community/:id", controllers.CommunityDetailHandler)

	v1.POST("/post", middleware.JWTAuthMiddleWare(), controllers.CreatePostHandler)

	//验证jwt机制
	v1.GET("/ping", middleware.JWTAuthMiddleWare(), func(context *gin.Context) {
		// 这里post man 模拟的 将token auth-token
		zap.L().Debug("ping", zap.String("ping-username", context.GetString("username")))
		controllers.ResponseSuccess(context, "pong")
	})

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return r
}
