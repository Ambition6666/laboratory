package middleware

import (
	"laboratory/pkg/utils"
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 验证用户是否登录的中间件
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.GetHeader("Authorization") //得到字串开头
		if t == "" || !strings.HasPrefix(t, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, "bearer解析失败")
			ctx.Abort()
			return
		}

		t = t[7:]                          //扔掉头部
		tk, c, r, e := utils.ParseToken(t) //c为claim结构体的实例
		if e != nil || !tk.Valid {
			ctx.JSON(http.StatusUnauthorized, "token解析失败")
			ctx.Abort() //中间件不通过
			return
		}
		//查找用户
		//存储用户信息
		ctx.Set("id", c)
		ctx.Set("role", r)
		ctx.Next()
	}
}
