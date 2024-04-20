package dao

import (
	"laboratory/model"
	"laboratory/sql"
)

// 创建老师
func CreateTeacher(t *model.Teacher) error {
	db := sql.GetMySQLDB()
	return db.Create(t).Error
}

// 通过邮箱获取用户信息
func GetInfoByEmailT(em string) *model.Teacher{
	db := sql.GetMySQLDB()	
	t := new(model.Teacher)
	db.Where("email = ?", em).Find(t)
	return t
}

// 修改老师信息
func UpdateTeacherINFO(t *model.Teacher) error{
	db := sql.GetMySQLDB()
	return db.Save(t).Error
}