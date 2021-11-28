package middleware

import (
	"go_web_app/controllers"
	"go_web_app/pkg/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuthMiddleWare() func(context *gin.Context) {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("auth-token")
		// 注意abort的写法 只要发现 token不合法 则直接abort 不会传到真正的路由那里
		if token == "" {
			controllers.ResponseError(context, controllers.CodeTokenIsEmpty)
			context.Abort()
			return
		}
		parseToken, err := jwt.ParseToken(token)
		if err != nil {
			controllers.ResponseError(context, controllers.CodeTokenInvalid)
			context.Abort()
			return
		}

		// 当token合法的时候 我们可以把对应的关键信息 取出来放到context里面，则对应的路由就可以直接通过get的方法 取出关键信息了
		zap.L().Debug("token parse", zap.String("username", parseToken.UserName))
		context.Set(controllers.ContextUserNameKey, parseToken.UserName)
		context.Set(controllers.ContextUserIdKey, parseToken.UserId)
		context.Next()
	}
}
