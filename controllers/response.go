package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	// omitempty 代表 data 为null 就不返回了
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}
