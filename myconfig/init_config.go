package myconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitGConfig() error {
	// 预加载环境变量
	viper.AutomaticEnv()
	// 获取一个viper实例
	vconfig := viper.New()
	// 设置配置文件的路径
	vconfig.SetConfigName("config")
	vconfig.AddConfigPath("../..")
	vconfig.SetConfigType("yaml")
	// 读取配置文件
	err := vconfig.ReadInConfig()
	if err != nil {
		_ = fmt.Errorf("service error: %s\n", err)
		return err
	}
	err = vconfig.Unmarshal(GConfig)
	if err != nil {
		_ = fmt.Errorf("service error: %s\n", err)
		return err
	}
	fmt.Println(GConfig)
	return nil
}
