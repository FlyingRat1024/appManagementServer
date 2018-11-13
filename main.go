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
	"androidappServer/api/material"
	"fmt"
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
import (
	"androidappServer/api/Login"
	"androidappServer/config"
)

func main() {
	config.InitLogger()
	cfg := config.InitConfig()
	serverAddr := fmt.Sprintf("%s:%d", cfg.API.Host, cfg.API.Port)
	router := gin.Default()
	router.Use(cors.Default())
	// app
	loginRouter := router.Group("/app")
	loginRouter.Handle("POST", "/login", Login.LoginHandler)
	// material
	materialRouter := router.Group("/material")
	receive := materialRouter.Group("/receive")
	receive.Handle("POST", "/write_table", material.WriteReceiveHandler)
	//

	//
	err := router.Run(serverAddr)
	if err != nil {
		logger.Error("run server error!")
	}
}
