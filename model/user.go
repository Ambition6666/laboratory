package model

 // user
 // 用户基本模型
type User struct {
	Email    string `json:"email"`   // 邮箱   
	Name     string `json:"name"`    // 姓名
	Password string `json:"_"`// 密码
	Phone    string `json:"phone"`   // 手机号
}