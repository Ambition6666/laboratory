package request

type LoginInfo struct{
	Email string // 可用学号或者邮箱登录
	Pwd string  // 密码	
}

type LoginInfoByAuth struct{
	Email string // 可用学号或者邮箱登录
	AuthCode string  // 验证码	
}


