package util

import (
	"gitee.com/nichanghao/gdmin/global"
	_ "gitee.com/nichanghao/gdmin/initialize"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	jwtCfg     = &global.Config.JWT
	signingKey = []byte(jwtCfg.SigningKey)
	JWT        = &JWTUtils{}
)

type JWTUtils struct {
}

func (jwtUtils *JWTUtils) GenerateToken(claims jwt.Claims) (string, error) {
	registeredClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtCfg.ExpiresTime) * time.Second)), // 过期时间
		Issuer:    jwtCfg.Issuer,                                                                       // 签名的发行者
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims).SignedString(signingKey)

}
