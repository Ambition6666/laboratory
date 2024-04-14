package utils

import (
	"laboratory/pkg/enum"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// -------------------------------------jwt生成token加密------------------------------------------------
type Claim struct {
	ID   int64
	jwt.RegisteredClaims
} //创建用户登录标签

// 得到token
func GetToken(id int64, role int) (string, error) {
	a := Claim{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), //token有效时间
			Issuer:    "zty",                                                   //签发人
		},
	} //获取claim实例
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a) //获取token
	return token.SignedString(enum.MSK)                   //返回加密串
}

// 解析token
func ParseToken(token string) (*jwt.Token, int64, error) {
	claim := &Claim{}
	t, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return enum.MSK, nil
	}) //接收前端发来加密字段
	return t, claim.ID, err
}

// ----------------------------------------使用sha256加密密码-----------------------------------------
func Encrypt(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password + enum.SALT)) //密码与盐自定义组合
	res := hex.EncodeToString(hash.Sum(nil))
	return res
}
