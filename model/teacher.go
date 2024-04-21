package model

import "gorm.io/gorm"

type Teacher struct {
	UINFO User `json:"basic" gorm:"embedded"` // 用户基本信息
	IsOK  bool `json:"isOk" gorm:"-"`
}

// 教师的构造器
func NewTeacher(email string, phone string, name string, pwd string) *Teacher {
	return &Teacher{
		UINFO: *NewUser(email, name, pwd, phone, 1),
	}
}

// 教师的信息是否合理
func (t *Teacher) IsInfoOK() bool {
	return t.UINFO.IsInfoOK()
}

// 教师信息是否完整
func (t *Teacher) AfterFind(tx *gorm.DB) (err error) {
	t.IsOK = t.IsInfoOK()
	return
}
