package user

import (
	"laboratory/internal/dao"
	"laboratory/log"
	"laboratory/model"
	"laboratory/vo/request"

	"net/http"
)

// 创建老师用户
func CreateTeacher(info *request.RegisteTeacherInfo) (int, string) {

	re, err := dao.GetUserRegister(info.Email)

	if !(err != nil || re == "1") {
		return http.StatusBadRequest, "已有该用户"
	}

	t := model.NewTeacher(info.Email, info.Phone, info.Name, info.Pwd)
	if !t.IsInfoOK() {
		return http.StatusBadRequest, "教师信息有误"
	}

	err = dao.CreateTeacher(t)

	if err != nil {
		log.SugarLogger.Error(err)
		return http.StatusBadRequest, "注册失败"
	}

	dao.SetUserRegister(info.Email)

	return http.StatusOK, "创建成功"

}

func GetTeacherINFO(id uint) (int, string, any) {
	s := new(model.Teacher)
	err := dao.GetUserInfoByID(id, s)

	if err != nil {
		log.SugarLogger.Error("获取教师信息错误", err)
		return http.StatusBadRequest, "获取失败", nil
	}

	return http.StatusOK, "获取成功", s
}

func UpdateTeacherINFO(id uint, data *request.UpdateTeaINFO) (int, string) {
	s := new(model.Teacher)
	err := dao.GetUserInfoByID(id, s)

	if err != nil {
		log.SugarLogger.Error("获取教师信息错误", err)
		return http.StatusBadRequest, "修改失败"
	}

	s.UINFO.Phone = data.Phone
	s.UINFO.Name = data.Name

	if s.IsInfoOK() {
		err = dao.UpdateTeacherINFO(s)
		if err != nil {
			log.SugarLogger.Error("更新教师信息错误", err)
			return http.StatusBadRequest, "修改失败"
		}
		return http.StatusOK, "修改成功"
	}

	return http.StatusBadRequest, "信息不符合规范"
}
