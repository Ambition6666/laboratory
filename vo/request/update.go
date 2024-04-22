package request

type UpdateStuINFO struct {
	Name    string `json:"name"`
	SID     string `json:"sid"`
	Phone   string `json:"phone"`
	Academy string `json:"academy"`
	Class   string `json:"class"`
}

type UpdateTeaINFO struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Academy string `json:"academy"`
}

type UpdateUserPWD struct {
	Email    string `json:"email"`
	AuthCode string `json:"auth_code"`
	PassWord string `json:"pwd"`
}
