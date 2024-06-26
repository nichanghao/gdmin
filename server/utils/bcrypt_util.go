package utils

import (
	"golang.org/x/crypto/bcrypt"
)

var (
	BCRYPT = &bcryptUtil{}
)

type bcryptUtil struct {
}

// HashPassword 对密码进行哈希处理
func (*bcryptUtil) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 验证密码是否匹配
func (*bcryptUtil) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
