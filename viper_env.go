package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	Production = iota
	Development
	Test
)

func envViper() *viper.Viper {
	vp := viper.New()
	vp.SetConfigType("yml")
	viperSearchPath(vp)

	switch {
	case Prod():
		vp.SetConfigName("production")
	case T():
		vp.SetConfigName("test")
	default:
		vp.SetConfigName("development")
	}
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintln("ViperEnv 加载错误", err))
	}
	return vp
}

// 设置当前应用的环境
func SetEnv(env int) {
	switch env {
	case Production:
		viper.Set(enviroment, "production")
	case Test:
		viper.Set(enviroment, "test")
	case Development:
		viper.Set(enviroment, "development")
	default:
		viper.Set(enviroment, "development")
	}
}

// 返回当前的应用的环境
func Enviroment() string {
	return viper.GetString("enviroment")
}

// 产品环境
func Prod() bool {
	return strings.EqualFold(Enviroment(), "production")
}

// 开发环境
func Dev() bool {
	return strings.EqualFold(Enviroment(), "development")
}

// 测试环境
func T() bool {
	return strings.EqualFold(Enviroment(), "test")
}
