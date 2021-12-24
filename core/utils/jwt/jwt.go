package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(uid string, role int, expireDuration time.Duration) (string, error) {

	expire := time.Now().Add(expireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, util.LoginClaims{
		Uid:  uid,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(util.SecretKey))
}
