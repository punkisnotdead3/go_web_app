package controllers

import (
	"fmt"
	"go_web_app/logic"
	"go_web_app/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamRegister)
	// 这里只能校验下 是否是标准的json格式 之类的 比较简单
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("RegisterHandler with invalid param", zap.Error(err))
		// 因为有的错误 比如json格式不对的错误 是不属于validator错误的 自然无法翻译，所以这里要做类型判断
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
		}
		return
	}
	// 进行参数校验
	//if p.RePassword != p.Password {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "密码和确认密码不同",
	//	})
	//	return
	//}
	fmt.Println(p)
	// 业务处理
	logic.Register(p)
	// 返回响应
	c.JSON(http.StatusOK, "register success")
}
