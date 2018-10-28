package config

import "github.com/donnie4w/go-logger/logger"

func InitLogger() {
	logger.SetRollingDaily("./log", "appServer.log")
	logger.SetLevel(logger.INFO)
}
