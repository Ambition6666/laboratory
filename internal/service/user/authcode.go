package user

import (
	"fmt"
	"laboratory/config"
	"laboratory/internal/dao"
	"laboratory/log"
	"laboratory/pkg/utils"
	"net/http"
	"net/smtp"

	e "github.com/jordan-wright/email"
)

// 发送验证码
func Send(em string) (int, string) {
	err := SendAuthCode(em)
	if err != nil {
		log.SugarLogger.Error("发送邮箱失败:", err)
		return http.StatusInternalServerError, "发送失败"
	}
	return http.StatusOK, "发送成功"
}

// ---------------------发送验证码----------------------------------
func SendAuthCode(to string) error {
	code, err := CreateAuthCode(to)
	if err != nil {
		return err
	}
	subject := "【实验室预约平台】邮箱验证"
	html := fmt.Sprintf(`<div style="text-align: center;">
		<h2 style="color: #333;">欢迎使用，你的验证码为：</h2>
		<h1 style="margin: 1.2em 0;">%s</h1>
		<p style="font-size: 12px; color: #666;">请在5分钟内完成验证，过期失效，请勿告知他人，以防个人信息泄露</p>
	</div>`, code)
	em := e.NewEmail()
	data := config.GetEmailInfo()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = data.From

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{to}

	// 设置主题
	em.Subject = subject

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.HTML = []byte(html)
	//fmt.Println(config.EmailAddr, config.Email, config.EmailAuth, config.EmailHost, config.EmailFrom)
	//设置服务器相关的配置
	err = em.Send(data.Addr, smtp.PlainAuth("", data.Email, data.Auth, data.Host))
	return err
}

func CreateAuthCode(em string) (string, error) {
	code := fmt.Sprintf("%d", utils.RandNum(900000)+100000)
	err := dao.SetAuthCode(em, code)
	return code, err
}

// 校验验证码
func IdentifyCode(em string, authCode string) (int, string) {
	res, err := dao.GetAuthCode(em)
	if err != nil {
		log.SugarLogger.Debug("获取验证码失败:", err)
		return http.StatusUnauthorized, "验证码失效"
	}
	if res != authCode {
		return http.StatusUnauthorized, "验证码错误"
	}
	return http.StatusOK, "验证成功"
}
