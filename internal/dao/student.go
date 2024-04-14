package dao

import (
	"laboratory/model"
	"laboratory/sql"
)

// 创建学生
func CreateStudent(s *model.Student) error {
	db := sql.GetMySQLDB()
	return db.Create(s).Error
}

// 通过邮箱获取用户信息
func GetInfoByEmailS(em string) *model.Student {
	db := sql.GetMySQLDB()
	s := new(model.Student)
	db.Where("email = ?", em).Find(s)
	return s
}
