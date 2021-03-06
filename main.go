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
	"androidappServer/api/message"
	"androidappServer/api/permission"
	"androidappServer/api/user"
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
	loginRouter.POST("/login", login.LoginHandler)
	loginRouter.GET("/permission", permission.PermissionHandler)
	// material
	materialRouter := router.Group("/material")
	materialRouter.GET("/", material.MaterialListHandler)
	materialRouter.POST("/create", material.CreateMaterialHandler)
	materialRouter.POST("/apply/write_table", material.WriteApplyTableHandler)
	materialRouter.GET("/apply", material.ApplyListhandler)
	materialRouter.GET("/apply/detail", material.ApplyDetailHandler)
	materialRouter.POST("/apply/verify", material.ApplyVerifyHandler)
	materialRouter.POST("/receive/write_table", material.WriteReceiveHandler)
	materialRouter.GET("/receive", material.ReceiveListhandler)
	materialRouter.GET("/receive/detail", material.ReceiveDetailHandler)
	materialRouter.POST("/receive/verify", material.ReceiveVerifyHandler)
	materialRouter.POST("/back/write_table", material.WriteBackTableHandler)
	materialRouter.POST("/waste_back/write_table", material.WriteWasteBackTableHandler)
	materialRouter.POST("/check/write_table", material.WriteCheckTableHandler)
	// warehouse
	warehouseRouter := router.Group("/warehouse")
	warehouseRouter.POST("/in", warehouse.WriteInWarehouseHandler)
	warehouseRouter.GET("/in", warehouse.InWarehouseListHandler)
	warehouseRouter.GET("/in/detail", warehouse.InWarehouseDetailHandler)
	warehouseRouter.POST("/out", warehouse.WriteOutWarehouseHandler)
	warehouseRouter.GET("/out", warehouse.OutWarehouseListHandler)
	warehouseRouter.POST("/out/confirm", warehouse.ConfirmOutWarehouseHandler)
	warehouseRouter.GET("/out/detail", warehouse.OutWarehouseDetailHandler)
	// user
	userRouter := router.Group("/user")
	userRouter.GET("/info", user.GetUserInfo)
	userRouter.POST("/modify/password", user.ModifyPassword)
	//message
	messageRouter := router.Group("/message")
	messageRouter.GET("/", message.MaterialWarningHandler)

	err := router.Run(serverAddr)
	if err != nil {
		logger.Error("run server error!")
	}
}
