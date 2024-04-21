package model

import (
	"laboratory/log"
	"regexp"

	"gorm.io/gorm"
)

// student
// 学生
type Student struct {
	UINFO   User   `json:"basic" gorm:"embedded"` // 用户基本信息
	SID     string `json:"sid" gorm:"not null" excel:"h"`   // 学号
	Academy string `json:"academy" excel:"h"`               // 学院
	Class   string `json:"class" excel:"h"`                 // 班级信息，例如计算机22-3
	IsOK    bool   `json:"isOk" gorm:"-"`         // 判断信息是否完整
}

// 学生的构造器
func NewStudent(email string, name string, pwd string, sid string) *Student {
	return &Student{
		UINFO:   *NewUser(email, name, pwd, "10000000000", 0),
		SID:     sid,
		Academy: "",
		Class:   "",
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

// 判断学生信息是否完整
func (s *Student) IsALLINFO() bool {
	return s.IsInfoOK() && s.Class != "" && s.Academy != ""
}

func (s *Student) AfterFind(tx *gorm.DB) (err error) {
	s.IsOK = s.IsALLINFO()
	return
}
