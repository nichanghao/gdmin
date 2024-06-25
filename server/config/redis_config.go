package config

type Redis struct {
	Addr           string
	Password       string
	DB             int
	MaxActiveConns int `mapstructure:"max-active-conns"` // 最大连接数
	MinIdleConns   int `mapstructure:"min-idle-conns"`   // 最小空闲连接数
	MaxIdleConns   int `mapstructure:"max-idle-conns"`   // 最大空闲连接数
}
