package config

type Gin struct {
	Address string // 服务监听地址
	Mode    string // gin运行模式 debug/test/release
}
