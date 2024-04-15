package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 校验是否有老师的权限
func IfTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetInt("role")
		if role < 1 {
			c.JSON(http.StatusForbidden, "权限不足")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// 校验是否有系统管理员的权限
func IfAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetInt("role")
		if role < 2 {
			c.JSON(http.StatusForbidden, "权限不足")
			c.Abort()
		} else {
			c.Next()
		}
	}
}
