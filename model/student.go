package model

import (
	"laboratory/log"
	"regexp"

	"gorm.io/gorm"
)

// student
// 学生
type Student struct {
	gorm.Model
	UINFO   *User // 用户基本信息
	SID     string `json:"sid"`     // 学号
	Academy string `json:"academy"` // 学院
	Class   string `json:"class"`   // 班级信息，例如计算机22-3
}

// 学生的构造器
func NewStudent(email string, name string, pwd string, sid string) *Student{
	return &Student{
		UINFO: NewUser(email, name, pwd, "10000000000"),
		SID:  sid,
		Academy: "",
		Class: "",
	}
}

// 判断是不是标准学号
func (s *Student) IsSID() bool {
	re, err := regexp.Compile(`[0-9]{11}`)

	if err != nil {
		log.SugarLogger.Errorf("正则表达式有误\n")
		return false
	}

	return re.MatchString(s.SID)
}

// 判断学生的信息是否ok
func (s *Student) IsInfoOK() bool {
	return s.UINFO.IsInfoOK() && s.IsSID()
}