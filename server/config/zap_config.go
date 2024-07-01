package config

type Zap struct {
	Path       string // 日志路径
	Filename   string // 日志文件名
	Level      string // 日志级别
	MaxSize    int    `mapstructure:"max-size"`    // 单个日志文件最大大小，单位：M
	MaxBackups int    `mapstructure:"max-backups"` // 日志文件最大备份数
	MaxAge     int    `mapstructure:"max-age"`     // 日志文件最长保留时间，单位：天
	Compress   bool   `mapstructure:"compress"`    // 是否压缩日志文件
}
