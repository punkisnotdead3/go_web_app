package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const ContextUserNameKey = "username"
const ContextUserIdKey = "userid"

var (
	ErrorNotLogin = errors.New("no login")
)

func getCurrentUserId(c *gin.Context) (int64, error) {
	userId, ok := c.Get(ContextUserIdKey)
	if !ok {
		return 0, ErrorNotLogin
	}
	return userId.(int64), nil
}

func getCurrentUserName(c *gin.Context) (string, error) {
	userName, ok := c.Get(ContextUserIdKey)
	if !ok {
		return "", ErrorNotLogin
	}
	return userName.(string), nil
}
