package util

import (
	"gopkg.in/ini.v1"
	"os"
)

type Config struct {
	Server `ini:"server"`
	Db     `ini:"db"`
	Logger `ini:"logger"`
}

// 服务器配置
type Server struct {
	Port string `ini:"port"`
}

// 数据库配置
type Db struct {
	Driver   string `ini:"driver"`
	Address  string `ini:"address"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	DataBase string `ini:"dataBase"`
	Charset  string `ini:"charset"`
	Port     string `ini:"port"`
}

// 日志配置
type Logger struct {
	Loglevel string `ini:"loglevel"`
	LogPath  string `ini:"logPath"`
}

// 配置文件转换
func GetConfig() (*Config, error) {
	var path string
	env := os.Getenv("GIN_MODE")
	switch env {
	case "release":
		path = "./config/config.ini"
	default:
		path = "./config/config-debug.ini"
	}
	config := new(Config)
	cfg, err := ini.LoadSources(ini.LoadOptions{
		IgnoreContinuation: true,
	}, path)

	if err != nil {
		return nil, err
	}

	err = cfg.MapTo(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
