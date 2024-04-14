package request

// 老师注册
type RegisteTeacherInfo struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Pwd string `json:"pwd"`
	AuthCode string `json:"auth_code"`
	Phone string `json:"phone"`
}

// 学生注册
type RegisterStudentInfo struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Pwd string `json:"pwd"`
	AuthCode string `json:"auth_code"`
	SID  string `json:"sid"`
}