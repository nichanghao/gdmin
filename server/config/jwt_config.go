package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key"`  // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time"` // 过期时间，单位：秒
	Issuer      string `mapstructure:"issuer"`       // 签发者
}
