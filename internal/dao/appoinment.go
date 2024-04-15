package dao

import (
	"laboratory/model"
	"laboratory/sql"
)

// 创建预约信息
func CreateAppointment(a *model.Appointment) error {
	db := sql.GetMySQLDB()
	return db.Create(a).Error
}
