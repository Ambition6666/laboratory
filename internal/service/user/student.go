package user

import (
	"laboratory/internal/dao"
	"laboratory/log"
	"laboratory/model"
	"laboratory/vo/request"
	"net/http"
)

// 创建老师用户
func CreateStudent(info *request.RegisterStudentInfo) (int, string) {

	re, err := dao.GetUserRegister(info.Email)

	if !(err != nil || re == "1") {
		return http.StatusBadRequest, "已有该用户"
	}

	s := model.NewStudent(info.Email, info.Name, info.Pwd, info.SID)
	if !s.IsInfoOK() {
		return http.StatusBadRequest, "学生信息有误"
	}

	err = dao.CreateStudent(s)

	if err != nil {
		log.SugarLogger.Error(err)
		return http.StatusBadRequest, "注册失败"
	}

	dao.SetUserRegister(info.Email)

	return http.StatusOK, "创建成功"

}

func GetStudentINFO(id uint) (int, string, any) {
	s := new(model.Student)
	err := dao.GetUserInfoByID(id, s)

	if err != nil {
		log.SugarLogger.Error("获取学生信息错误", err)
		return http.StatusBadRequest, "获取失败", nil
	}

	return http.StatusOK, "获取成功", s
}

func UpdateStudentINFO(id uint, data *request.UpdateStuINFO) (int, string) {
	s := new(model.Student)
	err := dao.GetUserInfoByID(id, s)

	if err != nil {
		log.SugarLogger.Error("获取学生信息错误", err)
		return http.StatusBadRequest, "修改失败"
	}

	s.SID = data.SID
	s.Academy = data.Academy
	s.Class = data.Class
	s.UINFO.Phone = data.Phone
	s.UINFO.Name = data.Name

	if s.IsInfoOK() {
		err = dao.UpdateStudentINFO(s)
		if err != nil {
			log.SugarLogger.Error("更新学生信息错误", err)
			return http.StatusBadRequest, "修改失败"
		}
		return http.StatusOK, "修改成功"
	}

	return http.StatusBadRequest, "信息不符合规范"
}
