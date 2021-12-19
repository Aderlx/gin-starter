package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	// 初始化 配置文件
	var v = viper.New()
	//设置配置文件的名字
	v.SetConfigName("dev.yaml")

	//添加配置文件所在的路径
	v.AddConfigPath("./config/envs")

	//设置配置文件类型，可选
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	return v
}
