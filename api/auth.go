package api

import (
	"laboratory/internal/service/user"
	"laboratory/vo/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取验证码
func GetAuthCode(c *gin.Context) {
	em := c.Query("email")
	code, msg := user.Send(em)
	c.JSON(http.StatusOK, response.Common{
		Code: code,
		Msg:  msg,
	})
}
