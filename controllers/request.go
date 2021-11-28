package controllers

import (
	"errors"
	"go_web_app/middleware"

	"github.com/gin-gonic/gin"
)

var (
	ErrorNotLogin = errors.New("no login")
)

func getCurrentUserId(c *gin.Context) (int64, error) {
	userId, ok := c.Get(middleware.ContextUserIdKey)
	if !ok {
		return 0, ErrorNotLogin
	}
	return userId.(int64), nil
}

func getCurrentUserName(c *gin.Context) (string, error) {
	userName, ok := c.Get(middleware.ContextUserIdKey)
	if !ok {
		return "", ErrorNotLogin
	}
	return userName.(string), nil
}
