package utils

import (
	"errors"
	"gitee.com/nichanghao/gdmin/global"
	"gitee.com/nichanghao/gdmin/model/common"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	JWT = &jwtUtil{}
)

type jwtUtil struct {
}

func (jwtUtils *jwtUtil) GenerateToken(userClaims *common.UserClaims) (string, error) {
	jwtCfg := &global.Config.JWT

	claims := common.JWTClaims{
		UserClaims: *userClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtCfg.ExpiresTime) * time.Second)), // 过期时间
			Issuer:    jwtCfg.Issuer,                                                                       // 签名的发行者
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtCfg.SigningKey))

}

func (jwtUtils *jwtUtil) ValidateToken(tokenStr string) (*common.UserClaims, error) {

	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac
	token, err := jwt.ParseWithClaims(tokenStr, &common.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.SigningKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, common.ErrTokenExpired
		}
		return nil, err
	}

	if jwtClaims, ok := token.Claims.(*common.JWTClaims); ok && token.Valid {
		return &jwtClaims.UserClaims, nil
	}

	return nil, errors.New("invalid token")

}
