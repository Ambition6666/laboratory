package user

import (
	"laboratory/internal/dao"
	"laboratory/log"
	"laboratory/model"
	"laboratory/pkg/utils"
	"net/http"
)

// 登录
func Login(em string, pwd string) (int, string) {
	u := dao.GetInfoByEmail(em)
	if u.ID == 0 {
		return http.StatusUnauthorized, "用户不存在"
	}
	if utils.Encrypt(pwd) != u.PassWord {
		return http.StatusUnauthorized, "密码错误"
	}

	token, err := utils.GetToken(u.ID, u.Role)

	if err != nil {
		log.SugarLogger.Error("token",err)
		return http.StatusInternalServerError, "获取token失败"
	}
	return http.StatusOK, token
}

// 用户通过验证码登录
func LoginByAuthCode(em string, authCode string) (int, string) {
	u := dao.GetInfoByEmail(em)
	if u.ID == 0 {
		return http.StatusUnauthorized, "用户不存在"
	}
	code, msg := IdentifyCode(em, authCode)

	if code != http.StatusOK {
		return code, msg
	}

	token, err := utils.GetToken(u.ID, u.Role)

	if err != nil {
		log.SugarLogger.Error("token",err)
		return http.StatusInternalServerError, "获取token失败"
	}

	return http.StatusOK, token
}


// 用户更改密码
func UpdateUserINFO( em string, authcode string, pwd string) (int, string) {
	u := dao.GetInfoByEmail(em)
	if u.ID == 0 {
		return http.StatusUnauthorized, "用户不存在"
	}
	code, msg := IdentifyCode(em, authcode)

	if code != http.StatusOK {
		return code, msg
	}

	switch u.Role {
	case 0:
		stu := new (model.Student)
		dao.GetUserInfoByID(u.ID, stu)
		stu.UINFO.PassWord = utils.Encrypt(pwd)
		err := dao.ChangeUserPassWord(stu)
		if err != nil {
			log.SugarLogger.Error("用户更改密码错误", err)
			return http.StatusInternalServerError, "更改失败"
		}
	case 1:
		tea := new (model.Teacher)
		dao.GetUserInfoByID(u.ID, tea)
		tea.UINFO.PassWord = utils.Encrypt(pwd)
		err := dao.ChangeUserPassWord(tea)
		if err != nil {
			log.SugarLogger.Error("用户更改密码错误", err)
			return http.StatusInternalServerError, "更改失败"
		}
	}

	return http.StatusOK, "更改成功"
}