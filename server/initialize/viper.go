package initialize

import (
	"fmt"
	"gitee.com/nichanghao/gdmin/global"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

func InitViper() {
	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("toml")
	// 添加配置文件路径
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 将配置绑定到结构体
	if err := viper.Unmarshal(&global.Config, func(dc *mapstructure.DecoderConfig) {}); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		// 重新读取配置并绑定到结构体
		if err := viper.Unmarshal(&global.Config, func(dc *mapstructure.DecoderConfig) {}); err != nil {
			log.Printf("Unable to decode into struct, %v", err)
		}
	})

}
