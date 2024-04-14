package api

import (
	"laboratory/internal/service/user"
	"laboratory/log"
	"laboratory/vo/request"
	"laboratory/vo/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 通过密码登录
func LoginByPwd(c *gin.Context) {
	info := new(request.LoginInfo)
	err := c.Bind(info)
	if err != nil {
		log.SugarLogger.Error("绑定数据失败:", err)
		c.JSON(http.StatusOK, response.Common{
			Code: http.StatusBadRequest,
			Msg:  "数据样式错误",
		})
		return
	}
	code, data := user.Login(info.Email, info.Pwd)
	if code != http.StatusOK {
		c.JSON(http.StatusOK, response.CommonData{
			Code: code,
			Msg:  data,
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, response.CommonData{
			Code: code,
			Msg:  "登录成功",
			Data: data,
		})
	}
}

// 通过验证码登录
func LoginByAuthCode(c *gin.Context) {
	info := new(request.LoginInfoByAuth)
	err := c.Bind(info)
	if err != nil {
		log.SugarLogger.Error("绑定数据失败:", err)
		c.JSON(http.StatusOK, response.Common{
			Code: http.StatusBadRequest,
			Msg:  "数据样式错误",
		})
		return
	}
	code, data := user.LoginByAuthCode(info.Email, info.AuthCode)
	if code != http.StatusOK {
		c.JSON(http.StatusOK, response.CommonData{
			Code: code,
			Msg:  data,
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, response.CommonData{
			Code: code,
			Msg:  "登录成功",
			Data: data,
		})
	}
}
