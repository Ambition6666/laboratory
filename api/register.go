package api

import (
	"laboratory/internal/service/user"
	"laboratory/log"
	"laboratory/vo/request"
	"laboratory/vo/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 教师注册
func RegisterTeacher(c *gin.Context) {
	info := new(request.RegisteTeacherInfo)
	err := c.Bind(info)
	if err != nil {
		log.SugarLogger.Error("绑定数据失败:", err)
		c.JSON(http.StatusOK, response.Common{
			Code: http.StatusBadRequest,
			Msg:  "数据样式错误",
		})
		return
	}
	code, msg := user.IdentifyCode(info.Email, info.AuthCode)
	if code != http.StatusOK {
		c.JSON(http.StatusOK, response.Common{
			Code: code,
			Msg:  msg,
		})
		return
	}

	code, msg = user.CreateTeacher(info)
	c.JSON(http.StatusOK, response.Common{
		Code: code,
		Msg:  msg,
	})
}

// 学生注册
func RegisterStudent(c *gin.Context) {
	info := new(request.RegisterStudentInfo)
	err := c.Bind(info)
	if err != nil {
		log.SugarLogger.Error("绑定数据失败:", err)
		c.JSON(http.StatusOK, response.Common{
			Code: http.StatusBadRequest,
			Msg:  "数据样式错误",
		})
		return
	}
	code, msg := user.IdentifyCode(info.Email, info.AuthCode)
	if code != http.StatusOK {
		c.JSON(http.StatusOK, response.Common{
			Code: code,
			Msg:  msg,
		})
		return
	}

	code, msg = user.CreateStudent(info)
	c.JSON(http.StatusOK, response.Common{
		Code: code,
		Msg:  msg,
	})	
}