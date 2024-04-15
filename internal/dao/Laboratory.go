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
