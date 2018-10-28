package config

import (
	"os"
	"path"
	"path/filepath"
	"sync"
)

import (
	"github.com/BurntSushi/toml"
	"github.com/donnie4w/go-logger/logger"
)

type Config struct {
	DB  DatabaseServer `toml:"mysql"`
	API APIServer      `toml:"server"`
}

type DatabaseServer struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DB       string `toml:"db"`
}

type APIServer struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

var (
	cfg *Config
	once sync.Once
)


func InitConfig() *Config {
	once.Do(func() {
		rootPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Error("get root path failed")
		}
		configFilePath := path.Join(rootPath, "config.toml")
		_, err = toml.DecodeFile(configFilePath, &cfg)
		if err != nil {
			logger.Error("parse config file error!")
			os.Exit(1)
		}
	})
	return cfg
}
