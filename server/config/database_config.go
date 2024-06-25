package config

type Database struct {
	Driver string
	Mysql  Mysql
}

type Mysql struct {
	DSN           string
	TablePrefix   string `mapstructure:"table-prefix"`   // 表前缀
	SingularTable bool   `mapstructure:"singular-table"` // 是否使用单数表名
	MaxIdleCount  int    `mapstructure:"max-idle-count"` // 最大空闲连接数
	MaxOpenConns  int    `mapstructure:"max-open-conns"` // 最大打开连接数
}
