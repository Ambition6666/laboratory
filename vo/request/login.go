package request

type LoginInfo struct {
	Email string `json:"email"` // 可用学号或者邮箱登录
	Pwd   string `json:"pwd"`   // 密码
}

type LoginInfoByAuth struct {
	Email    string `json:"email"`     // 可用学号或者邮箱登录
	AuthCode string `json:"auth_code"` // 验证码
}
