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
	"androidappServer/api/warehouse"
	"fmt"
)

import (
	"github.com/donnie4w/go-logger/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
import (
	"androidappServer/api/login"
	"androidappServer/api/material"
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
	loginRouter.Handle("POST", "/login", login.LoginHandler)
	// material
	materialRouter := router.Group("/material")
	materialRouter.GET("/", material.MaterialListHandler)
	materialRouter.POST("/apply/write_table", material.WriteApplyTableHandler)
	materialRouter.GET("/apply/", material.ApplyListhandler)
	materialRouter.GET("/apply/detail", material.ApplyDetailHandler)
	materialRouter.POST("/receive/write_table", material.WriteReceiveHandler)
	materialRouter.GET("/receive/", material.ReceiveListhandler)
	materialRouter.GET("/receive/detail", material.ReceiveDetailHandler)
	materialRouter.POST("/back/write_table", material.WriteBackTableHandler)
	materialRouter.POST("/check/write_table", material.WriteCheckTableHandler)
	//
	repoRouter := router.Group("/warehouse")
	repoRouter.POST("/in/write_table", warehouse.WriteInWarehouseHandler)
	repoRouter.POST("/out/write_table", warehouse.WriteOutWarehouseHandler)

	err := router.Run(serverAddr)
	if err != nil {
		logger.Error("run server error!")
	}
}
