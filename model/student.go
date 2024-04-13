package model

import "gorm.io/gorm"

// student
// 学生
type Student struct {
	gorm.Model
	User
	SID string `json:"sid"` // 学号
	Academy string `json:"academy"`// 学院
	Class   string `json:"class"`  // 班级信息，例如计算机22-3
}