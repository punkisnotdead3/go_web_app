package controllers

import (
	"errors"
	"go_web_app/logic"
	"go_web_app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 点赞 点踩
func PostLikeHandler(c *gin.Context) {

	p := new(models.ParamLikeData)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("PostLikeHandler with invalid param", zap.Error(err))
		// 因为有的错误 比如json格式不对的错误 是不属于validator错误的 自然无法翻译，所以这里要做类型判断
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	id, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNoLogin)
		return
	}

	// 业务处理
	err = logic.PostLike(p, id)
	if err != nil {
		// 可以把
		if errors.Is(err, logic.ErrAleadyLike) || errors.Is(err, logic.ErrAleadyUnLike) {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, "success")
}
