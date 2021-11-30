package controllers

import (
	"go_web_app/logic"
	"strconv"

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

func CommunityDetailHandler(c *gin.Context) {
	communityIdStr := c.Param("id")
	communityId, err := strconv.ParseInt(communityIdStr, 10, 64)
	// 校验参数是否正确
	if err != nil {
		zap.L().Error("GetCommunityListDetail error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityById(communityId)
	if err != nil {
		zap.L().Error("GetCommunityListDetail error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
