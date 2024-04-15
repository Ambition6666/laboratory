package api

import (
	"laboratory/internal/service/booking"
	"laboratory/log"
	"laboratory/vo/request"
	"laboratory/vo/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 添加可预约实验室
func AddLaboratory(c *gin.Context) {
	info := new(request.LaboratoryInfo)
	err := c.Bind(info)
	if err != nil {
		log.SugarLogger.Error("绑定数据失败:", err)
		c.JSON(http.StatusOK, response.Common{
			Code: http.StatusBadRequest,
			Msg:  "数据样式错误",
		})
		return
	}

	tid := c.GetUint("id")

	code, msg := booking.PublishLaboratory(info.Date, info.Place, info.Raa, tid)
	c.JSON(http.StatusOK, response.Common{
		Code: code,
		Msg:  msg,
	})
}

// 学生预约实验室
func AddAppointment(c *gin.Context) {
	info := new(request.BookingInfo)
	err := c.Bind(info)
	if err != nil {
		log.SugarLogger.Error("绑定数据失败:", err)
		c.JSON(http.StatusOK, response.Common{
			Code: http.StatusBadRequest,
			Msg:  "数据样式错误",
		})
		return
	}
	sid := c.GetUint("id")
	code, msg := booking.BookingLaboratory(sid, info.LID, info.ProgramContent, info.Raa)
	c.JSON(http.StatusOK, response.Common{
		Code: code,
		Msg:  msg,
	})
}
