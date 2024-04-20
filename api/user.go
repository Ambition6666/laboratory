package api

import (
	"laboratory/internal/service/user"
	"laboratory/vo/request"
	"laboratory/vo/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户获取信息
func GetUserINFO(c *gin.Context) {
	role := c.GetInt("role")
	id := c.GetUint("id")
	var (
		code int
		data any
		msg  string
	)

	switch role {
	case 0:
		code, msg, data = user.GetStudentINFO(id)
	case 1:
		code, msg, data = user.GetTeacherINFO(id)
	default:
		code = http.StatusBadRequest
		data = nil
		msg = "参数无效"
	}

	c.JSON(http.StatusOK, response.CommonData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}


// 更新用户信息
func UpdateUserINFO(c *gin.Context) {
	role := c.GetInt("role")
	id := c.GetUint("id")
	var (
		code int
		msg  string
	)

	switch role {
	case 0:
		data := new(request.UpdateStuINFO)
		c.Bind(data)
		code, msg = user.UpdateStudentINFO(id, data)

	case 1:
		data := new(request.UpdateTeaINFO)
		c.Bind(data)
		code, msg = user.UpdateTeacherINFO(id, data)
		
	default:
		code = http.StatusBadRequest
		msg = "参数无效"
	}

	c.JSON(http.StatusOK, response.Common{
		Code: code,
		Msg:  msg,
	})
}