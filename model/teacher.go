package model

type Teacher struct {
	UINFO  User `gorm:"embedded"`// 用户基本信息
}

// 教师的构造器
func NewTeacher(email string, phone string, name string, pwd string) *Teacher {
	return  &Teacher{
		UINFO: *NewUser(email, name, pwd, phone, 1),
	}
}

// 教师的信息是否合理
func (t *Teacher) IsInfoOK() bool {
	return t.UINFO.IsInfoOK()
}

