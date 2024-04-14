package user

import (
	"laboratory/internal/dao"
	"laboratory/pkg/utils"
	"net/http"
)

// 登录
func Login(em string, pwd string) (int, string) {
	u := dao.GetInfoByEmailT(em)
	if u.ID == 0 {
		return http.StatusUnauthorized, "用户不存在"
	}
	if utils.Encrypt(pwd) != u.UINFO.PassWord {
		return http.StatusUnauthorized, "密码错误"
	}

	token, err := utils.GetToken(u.ID, u.UINFO.Role)

	if err != nil {
		return http.StatusInternalServerError, "获取token失败"
	}
	return http.StatusOK, token
}

// 用户通过验证码登录
func LoginByAuthCode(em string, authCode string) (int, string) {
	u := dao.GetInfoByEmailT(em)
	if u.ID == 0 {
		return http.StatusUnauthorized, "用户不存在"
	}
	code, msg := IdentifyCode(em, authCode)

	if code != http.StatusOK {
		return code, msg
	}

	token, err := utils.GetToken(u.ID, u.UINFO.Role)

	if err != nil {
		return http.StatusInternalServerError, "获取token失败"
	}

	return http.StatusOK, token
}
