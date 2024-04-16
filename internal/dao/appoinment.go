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

// 查询已经预约的实验室的人员信息
func SearchAppointment(lid string) ([]model.Appointment, error){
	db := sql.GetMySQLDB()
	list := make([]model.Appointment, 0)
	err := db.Where("l_id = ?", lid).Find(&list).Error
	return list, err
}

// 查询学生已经预约的实验室信息
func SearchStuAppointment(sid uint) ([]model.Appointment, error){
	db := sql.GetMySQLDB()
	list := make([]model.Appointment, 0)
	err := db.Where("s_id = ?", sid).Find(&list).Error
	return list, err
}