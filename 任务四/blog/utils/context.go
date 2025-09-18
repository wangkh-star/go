package utils

import (
	"blog/global"

	"github.com/gin-gonic/gin"
)

// 从 Context 获取用户信息
func GetUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		global.Log.Error("获取用户信息失败")
		return 0, false
	}
	return userID.(uint), true
}

func GetUsername(c *gin.Context) (string, bool) {
	username, exists := c.Get("username")
	if !exists {
		return "", false
	}
	return username.(string), true
}
