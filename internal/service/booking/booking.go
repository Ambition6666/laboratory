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

// 教师预约实验室
func BookingLaboratoryByTeacher(tid uint, t *model.TAppointment) (int, string){
	t.TID = tid
	err := dao.CreateTAppointment(t)
	if err != nil {
		log.SugarLogger.Error("创建教师预约信息错误", err)
		return http.StatusInternalServerError, "预约失败"
	} 

	return http.StatusOK, "预约成功"
}

// 教师查询已经预约的实验室
func SearchTeaAppointment(tid uint) (int, string, any) {
	list, err := dao.SearchTAppointment(tid)
	if err != nil {
		log.SugarLogger.Error("教师查询已预约实验室错误", err)
		return http.StatusBadRequest, "查询失败", nil
	}
	return http.StatusOK, "查询成功", list
}