package controllers

import (
	"go_web_app/logic"
	"go_web_app/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	postIdStr := c.Param("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	// 校验参数是否正确
	if err != nil {
		zap.L().Error("GetPostDetailHandler", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostDetail(postId)
	if err != nil {
		zap.L().Error("GetPostDetailHandler", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// 发帖
func CreatePostHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.Post)
	// 这里只能校验下 是否是标准的json格式 之类的 比较简单
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreatePostHandler with invalid param", zap.Error(err))
		// 因为有的错误 比如json格式不对的错误 是不属于validator错误的 自然无法翻译，所以这里要做类型判断
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 取不到userId 则提示需要登录
	authorId, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNoLogin)
		return
	}
	p.AuthorId = authorId
	msg, err := logic.CreatePost(p)
	zap.L().Info("CreatePostHandlerSuccess", zap.String("postId", msg))
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)

	// 创建帖子
	// 返回响应
}
