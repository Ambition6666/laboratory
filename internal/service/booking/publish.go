package booking

import (
	"laboratory/internal/dao"
	"laboratory/log"
	"laboratory/model"
	"net/http"
)

// 发布实验室开放名单
func PublishLaboratory(date string, place string, raa []string, tid uint) (int, string) {
	l := model.NewLaboratory(date, place, raa, tid)
	err := dao.CreateLaboratory(l)
	if err != nil {
		log.SugarLogger.Error("创建实验室开放名单错误`", err)
		return http.StatusInternalServerError, "创建失败"
	}

	return http.StatusOK, "创建成功"
}

// 查询发布实验室的报名名单
func SearchLaboratoryHuInfo(lid string) (int, string, any) {
	list, err := dao.SearchAppointment(lid)
	if err != nil {
		log.SugarLogger.Error("查询实验室人员名单错误", err)
		return http.StatusBadRequest, "查询失败", nil
	}
	return http.StatusOK, "查询成功", list
}
