package route

import (
	"go_web_app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}
	r := gin.New()
	// 最重要的就是这个日志库
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return r
}
