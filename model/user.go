package model

import (
	"laboratory/log"
	"laboratory/pkg/utils"
	"regexp"

	"gorm.io/gorm"
)

type USER interface {
	IsInfoOK() bool
}

// user
// 用户基本模型
type User struct {
	gorm.Model `json:"-"`
	Email      string `json:"email" gorm:"not null; unique; index"` // 邮箱
	Name       string `json:"name" excel:"h"`                                 // 姓名
	PassWord   string `json:"-"`                                    // 密码
	Phone      string `json:"phone"`                                // 手机号
	Role       int    `json:"rl"`                                   // 权限
}

// 用户构造器
func NewUser(email string, name string, password string, phone string, role int) *User {
	return &User{
		Email:    email,
		Name:     name,
		PassWord: utils.Encrypt(password),
		Phone:    phone,
		Role:     role,
	}
}

// 判断是不是一个email
// 为什么用反引号?
func (u *User) IsEmail() bool {
	re, err := regexp.Compile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)

	if err != nil {
		log.SugarLogger.Errorf("正则表达式有误\n")
		return false
	}

	return re.MatchString(u.Email)
}

// 判断是不是一个手机号码
func (u *User) IsPhone() bool {
	re, err := regexp.Compile(`[0-9]{11}`)

	if err != nil {
		log.SugarLogger.Errorf("正则表达式有误\n")
		return false
	}

	return re.MatchString(u.Phone)
}

// 判断名字长度不超过指定长度
func (u *User) IsNormalName() bool {
	return len(u.Name) > 0 && len(u.Name) < 20
}

// 判断用户信息是否全部符合标准
func (u *User) IsInfoOK() bool {
	return u.IsEmail() && u.IsPhone() && u.IsNormalName()
}
