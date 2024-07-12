package common

import "github.com/golang-jwt/jwt/v5"

// JWTClaims jwt claims
type JWTClaims struct {
	UserClaims
	jwt.RegisteredClaims
}

// UserClaims user claims
type UserClaims struct {
	ID       uint64
	Username string
	NickName string
}
