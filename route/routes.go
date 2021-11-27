package route

import (
	"go_web_app/controllers"
	"go_web_app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 最重要的就是这个日志库
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册
	r.POST("/register", controllers.RegisterHandler)
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return r
}
