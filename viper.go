package config

import (
	"github.com/spf13/viper"
)

const (
	enviroment = "enviroment"
)

func init() {
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
	return err
}

func Unmarshal(rawVal interface{}) error {
	return viper.Unmarshal(rawVal)
}

// viperSearchPath 设置 viper 搜索配置文件的路径
func viperSearchPath(vp *viper.Viper) {
	// 在项目根目录
	vp.AddConfigPath("./config")
	// 在 config 目录中(test)
	vp.AddConfigPath(".")
}
