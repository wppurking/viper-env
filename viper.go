package config

import (
	"time"

	//"github.com/asaskevich/govalidator"
	"github.com/spf13/viper"
)

const (
	enviroment = "enviroment"
)

func init() {
	// 重置整个系统的时区为 UTC, 避免系统间有时区的问题.
	// 如果是拥有 UI 的系统, 则建议使用本地时区为应用时区, 如果是后端应用, 统一使用 UTC 时区可避免
	// 应用在不同服务器上运行因为时区的问题发生的各种异常
	time.Local = time.UTC

	viper.SetConfigName("enviroment")
	viper.SetConfigType("yml")
	viperSearchPath(viper.GetViper())

	// Default Config
	viper.SetDefault(enviroment, "development")
	viper.SetDefault("api.port", ":3001")

	viper.SetEnvPrefix("ear")
	viper.AutomaticEnv()
}

// 加载整个应用的配置文件
func Load() error {
	// 1. 加载默认的配置文件
	// 2. 确定应用运行环境
	// 3. 加载对应环境的配置文件, 然后进行合并
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// 将新的环境变量的配置合并进入主配置文件
	ev := envViper()
	err = viper.MergeConfigMap(ev.AllSettings())
	if err != nil {
		return err
	}
	/*
		err = viper.Unmarshal(&C)
		if err != nil {
			return err
		}
		_, err = govalidator.ValidateStruct(C)
	*/
	return err
}

// viperSearchPath 设置 viper 搜索配置文件的路径
func viperSearchPath(vp *viper.Viper) {
	// 在项目根目录
	vp.AddConfigPath("./config")
	// 在 config 目录中(test)
	vp.AddConfigPath(".")
}
