package config

import (
	"go-devops-admin/pkg/helpers"
	"os"

	"github.com/spf13/cast"
	viperLib "github.com/spf13/viper"
)

// viper 库实例
var viper *viperLib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 加载到此数组, loadConfig动态加载配置
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 初始化viper库
	viper = viperLib.New()

	// 配置类型,支持"json","toml","yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	viper.SetConfigType("yml")

	// 环境变量配置文件查找路径, 和main.go同级
	viper.AddConfigPath(".")

	// 设置环境变量前缀，用于区分go系统变量
	viper.SetEnvPrefix("app_env")

	// 读取环境变量(支持flags)
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	// 加载环境变量
	LoadEnv(env)

	// 注册配置信息
	LoadConfig()
}

func LoadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func LoadEnv(envSuffix string) {
	// 默认加载.env文件, 如果参数传递--env=name，加载.env.name文件
	envPath := "settings"
	if len(envSuffix) > 0 {
		filepath := envPath + "-" + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			// 如果.env.name文件存在，则加载对应的配置文件，否则加载默认配置文件
			envPath = filepath
		}
	}

	// 加载env
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	// 监控 .env 配置文件，变更时重新加载
	viper.WatchConfig()
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return viper.Get(path)
}

// 获取配置
// 第一个参数path允许试用.式获取，如app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

// 获取String类型配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// 获取Int类型配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// 获取Float64类型配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// 获取Unit类型配置信息
func GetUnit(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// 获取bool类型配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// 获取stringMapString类型配置信息
func GetStringMapString(path string, defaultValue ...interface{}) map[string]string {
	return cast.ToStringMapString(internalGet(path, defaultValue...))
}
