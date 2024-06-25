package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key"`  // jwt签名
	ExpiresTime string `mapstructure:"expires-time"` // 过期时间
}
