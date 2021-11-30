package controllers

import (
	"go_web_app/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityHandler 板块
func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
