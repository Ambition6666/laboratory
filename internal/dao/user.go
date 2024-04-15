package dao

import (
	"context"
	"laboratory/model"
	"laboratory/sql"
	"time"
)

// 存储验证码
func SetAuthCode(em string, auth_code string) error {
	rdb := sql.GetRedisDB()
	return rdb.Set(context.Background(), "auth"+em, auth_code, 300 * time.Second).Err()
}

// 获取验证码
func GetAuthCode(em string) (string, error) {
	rdb := sql.GetRedisDB()
	return rdb.Get(context.Background(), "auth"+em).Result()
}

// 标记用户已经注册
func SetUserRegister(em string) error {
	rdb := sql.GetRedisDB()
	return rdb.Set(context.Background(), em, 1, 0).Err()
}

// 注销用户标记
func SetUserUnRegister(em string) error {
	rdb := sql.GetRedisDB()
	return rdb.Set(context.Background(), em, 0, 0).Err()	
}

// 获取用户是否注册的标记
func GetUserRegister(em string) (string, error) {
	rdb := sql.GetRedisDB()
	return rdb.Get(context.Background(), em).Result()
}

// 通过邮箱获取用户信息
func GetInfoByEmail(em string) *model.User {
	db := sql.GetMySQLDB()
	s := new(model.Student)
	db.Where("email = ?", em).Find(s)
	t := new(model.Teacher)
	if (s.UINFO.ID == 0) {
		db.Where("email = ?", em).Find(t)	
	} else {
		return &s.UINFO
	}
	return &t.UINFO
}