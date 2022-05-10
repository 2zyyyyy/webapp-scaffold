package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig() // 读取配置文件
	if err != nil {
		fmt.Println("viper.ReadInConfig failed, err:", err)
		return
	}
	viper.WatchConfig() // 热加载配置文件
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改!")
	})
	return
}
