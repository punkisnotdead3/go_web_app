package route

import (
	"go_web_app/controllers"
	"go_web_app/logger"
	"go_web_app/pkg/jwt"
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
	r.GET("/ping", func(context *gin.Context) {
		// 这里post man 模拟的 将token auth-token
		token := context.Request.Header.Get("auth-token")
		if token == "" {
			controllers.ResponseError(context, controllers.CodeTokenIsEmpty)
			return
		}
		parseToken, err := jwt.ParseToken(token)
		if err != nil {
			controllers.ResponseError(context, controllers.CodeTokenInvalid)
			return
		}

		zap.L().Debug("token parese", zap.String("username", parseToken.UserName))
		controllers.ResponseSuccess(context, "pong")
	})

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return r
}
