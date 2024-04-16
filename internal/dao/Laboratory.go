package dao

import (
	"laboratory/model"
	"laboratory/sql"
)

// 创建实验室列表
func CreateLaboratory(l *model.Laboratory) error {
	db := sql.GetMySQLDB()
	return db.Create(l).Error
}


// 查询实验室列表
func SearchLaboratory(date string) ([]model.Laboratory, error) {
	db := sql.GetMySQLDB()
	list := make([]model.Laboratory, 0)
	err := db.Where("date = ?", date).Find(&list).Error
	return list, err
}