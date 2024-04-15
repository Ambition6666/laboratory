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
