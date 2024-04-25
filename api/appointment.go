package api

import (
	"laboratory/internal/service/booking"
	"laboratory/internal/service/summary"
	"laboratory/log"
	"laboratory/model"
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

// 查询可预约的实验室
func SearchLaboratory(c *gin.Context) {
	date := c.Query("date")
	code, msg, data := booking.SearchLaboratory(date)
	c.JSON(http.StatusOK, response.CommonData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 学生查询已预约的实验室信息
func SearchAppointment(c *gin.Context) {
	sid := c.GetUint("id")
	code, msg, data := booking.SearchStuAppointment(sid)
	c.JSON(http.StatusOK, response.CommonData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 老师教询开放实验室已经预约的人员的名单
func SearchLaboratoryHuInfo(c *gin.Context) {
	lid := c.Query("lid")
	code, msg, data := booking.SearchLaboratoryHuInfo(lid)
	c.JSON(http.StatusOK, response.CommonData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 老师导出开放实验室已经预约的人员的名单
func ExportExcelHuInfo(c *gin.Context) {
	lid := c.Query("lid")
	code, msg, data := summary.GetBookingINFOExcel(lid)

	c.JSON(http.StatusOK, response.CommonData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 教师查询可预约的实验室
func SearchLaboratoryByTeacher(c *gin.Context) {
	tid := c.GetUint("id")
	code, msg, data := booking.SearchLaboratoryByTeacher(tid)
	c.JSON(http.StatusOK, response.CommonData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 教师申请实验室
func BookingLaboratoryByTeacher(c *gin.Context) {
	t := new(model.TAppointment)
	c.Bind(t)
	tid := c.GetUint("id")
	code, msg := booking.BookingLaboratoryByTeacher(tid, t)
	c.JSON(http.StatusOK, response.Common{
		Code: code,
		Msg:  msg,
	})
}

// 教师查询已经预约的实验室
func SearchMyBookingLaboratoryByTeacher(c *gin.Context) {
	tid := c.GetUint("id")
	code, msg, data := booking.SearchTeaAppointment(tid)
	c.JSON(http.StatusOK, response.CommonData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}