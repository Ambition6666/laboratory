package api

import (
	"laboratory/internal/service/user"
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
