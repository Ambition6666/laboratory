package biz

import (
	"laboratory/model"
	"laboratory/log"
	"testing"
)

func TestRexp(t *testing.T) {
	log.InitLogger()
	u := model.NewTeacher("1111111111@qq.com", "11111111111", "zty", "123456")
	if u.IsInfoOK() {
		t.Error("对的")
	} else {
		if !u.UINFO.IsEmail() {t.Error("邮箱")}
		if !u.UINFO.IsNormalName() {t.Error("名字")}
		if !u.UINFO.IsPhone() {t.Error("手机")}
	}
}
