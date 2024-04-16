package booking

import (
	"laboratory/internal/dao"
	"laboratory/log"
	"laboratory/model"
	"net/http"
)

// 预约实验室
func BookingLaboratory(sid uint, lid uint, pc string, raa []string) (int, string) {
	a := model.NewAppointment(sid, lid, pc, raa)
	err := dao.CreateAppointment(a)
	if err != nil {
		log.SugarLogger.Error("创建预约名单错误`", err)
		return http.StatusInternalServerError, "创建失败"
	}

	return http.StatusOK, "创建成功"
}

// 查询可预约的实验室
func SearchLaboratory(date string) (int, string, any) {
	list, err := dao.SearchLaboratory(date)
	if err != nil {
		log.SugarLogger.Error("查询预约实验室错误", err)
		return http.StatusBadRequest, "查询失败", nil
	}
	return http.StatusOK, "查询成功", list
}

// 查询已经预约的实验室
func SearchStuAppointment(sid uint) (int, string, any) {
	list, err := dao.SearchStuAppointment(sid)
	if err != nil {
		log.SugarLogger.Error("查询已预约实验室错误", err)
		return http.StatusBadRequest, "查询失败", nil
	}
	return http.StatusOK, "查询成功", list
}