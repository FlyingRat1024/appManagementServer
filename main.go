/*
Modification history
--------------------
2018/10/28, by Meng Yuhang created
*/
/*
 DESCRIPTION

*/
package main

import (
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
import (
	"androidappServer/api"
	"androidappServer/config"
)

func main() {
	config.InitLogger()
	cfg := config.InitConfig()
	serverAddr := fmt.Sprintf("%s:%d", cfg.API.Host, cfg.API.Port)
	router := gin.Default()
	router.Use(cors.Default())
	router.Handle("POST", "app/login", controller.LoginHandler)
	err := router.Run(serverAddr)
	if err != nil {
		logger.Error("run server error!")
	}
}
